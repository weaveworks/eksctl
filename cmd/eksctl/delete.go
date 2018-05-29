package main

import (
	"fmt"
	"os"

	"github.com/kubicorn/kubicorn/pkg/logger"
	"github.com/spf13/cobra"

	"github.com/weaveworks/eksctl/pkg/eks"
)

func deleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete",
		Run: func(c *cobra.Command, _ []string) {
			c.Help()
		},
	}

	cmd.AddCommand(deleteClusterCmd())

	return cmd
}

func deleteClusterCmd() *cobra.Command {
	cfg := &eks.Config{}

	cmd := &cobra.Command{
		Use: "cluster",
		Run: func(_ *cobra.Command, _ []string) {
			if err := doDeleteCluster(cfg); err != nil {
				logger.Critical(err.Error())
				os.Exit(1)
			}
		},
	}

	fs := cmd.Flags()

	fs.StringVarP(&cfg.ClusterName, "cluster-name", "n", "", "EKS cluster name (required)")
	fs.StringVarP(&cfg.Region, "region", "r", DEFAULT_EKS_REGION, "AWS region")

	return cmd
}

func doDeleteCluster(cfg *eks.Config) error {
	ctl := eks.New(cfg)

	if err := ctl.CheckAuth(); err != nil {
		return err
	}

	if cfg.ClusterName == "" {
		return fmt.Errorf("--cluster-name must be set")
	}

	logger.Info("deleting EKS cluster %q", cfg.ClusterName)

	if err := ctl.DeleteStackServiceRole(); err != nil {
		return err
	}

	if err := ctl.DeleteStackVPC(); err != nil {
		return err
	}

	if err := ctl.DeleteStackDefaultNodeGroup(); err != nil {
		return err
	}

	logger.Success("all EKS cluster %q resource will be deleted (if in doubt, check CloudForamtion console)", cfg.ClusterName)

	return nil
}
