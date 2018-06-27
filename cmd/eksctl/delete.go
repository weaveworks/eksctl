package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/kubicorn/kubicorn/pkg/logger"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	awseks "github.com/aws/aws-sdk-go/service/eks"
	"github.com/weaveworks/eksctl/pkg/eks"
	"github.com/weaveworks/eksctl/pkg/utils"
)

func deleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete resource(s)",
		Run: func(c *cobra.Command, _ []string) {
			c.Help()
		},
	}

	cmd.AddCommand(deleteClusterCmd())

	return cmd
}

func deleteClusterCmd() *cobra.Command {
	cfg := &eks.ClusterConfig{}

	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "Delete a cluster",
		Run: func(_ *cobra.Command, args []string) {
			if len(args) > 0 {
				cfg.ClusterName = strings.TrimSpace(args[0])
			}
			if err := doDeleteCluster(cfg); err != nil {
				logger.Critical(err.Error())
				os.Exit(1)
			}
		},
	}

	fs := cmd.Flags()

	fs.StringVarP(&cfg.ClusterName, "name", "n", "", "EKS cluster name (required)")
	fs.StringVarP(&cfg.Region, "region", "r", DEFAULT_EKS_REGION, "AWS region")
	fs.StringVarP(&cfg.Profile, "profile", "p", "", "AWS profile to use. If provided, this overrides the AWS_PROFILE environment variable")

	return cmd
}

func doDeleteCluster(cfg *eks.ClusterConfig) error {
	ctl := eks.New(cfg)

	if err := ctl.CheckAuth(); err != nil {
		return err
	}

	if cfg.ClusterName == "" {
		return fmt.Errorf("--name must be set")
	}

	logger.Info("deleting EKS cluster %q", cfg.ClusterName)

	if err := ctl.DeleteControlPlane(); err != nil {
		if !utils.HasAwsErrorCode(err, awseks.ErrCodeResourceNotFoundException) {
			return err
		} else {
			logger.Info("EKS cluster %q does not exist.", cfg.ClusterName)
		}
	}

	if err := ctl.DeleteStackServiceRole(); err != nil {
		if !utils.HasAwsErrorCode(err, cloudformation.ErrCodeStackInstanceNotFoundException) {
			return err
		} else {
			logger.Info("service role stack does not exist.")
		}
	}

	if err := ctl.DeleteStackDefaultNodeGroup(); err != nil {
		if !utils.HasAwsErrorCode(err, awseks.ErrCodeResourceNotFoundException) {
			return err
		} else {
			logger.Info("node group stack does not exist.")
		}
	}

	if err := ctl.DeleteStackVPC(); err != nil {
		if !utils.HasAwsErrorCode(err, awseks.ErrCodeResourceNotFoundException) {
			return err
		} else {
			logger.Info("VPC stack does not exist.")
		}
	}

	ctl.MaybeDeletePublicSSHKey()

	utils.MaybeDeleteConfig(cfg.ClusterName)

	logger.Success("all EKS cluster %q resource will be deleted (if in doubt, check CloudFormation console)", cfg.ClusterName)

	return nil
}
