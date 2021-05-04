package fargate

import (
	"fmt"
	"strings"

	"github.com/kris-nova/logger"
	"github.com/pkg/errors"
	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/cfn/builder"
	"github.com/weaveworks/eksctl/pkg/cfn/manager"
	"github.com/weaveworks/eksctl/pkg/utils/tasks"
)

type createFargateStackTask struct {
	cfg          *api.ClusterConfig
	provider     api.ClusterProvider
	stackManager manager.StackManager
}

func (t *createFargateStackTask) Describe() string { return "create fargate IAM stacK" }

func makeClusterStackName(clusterName string, disableStackPrefix bool) string {
	stackName := fmt.Sprintf("%s-fargate", clusterName)
	if !disableStackPrefix {
		stackName = "eksctl-" + stackName
	}
	return strings.Replace(stackName, "_", "-", -1)
}

func (t *createFargateStackTask) Do(errs chan error) error {
	rs := builder.NewFargateResourceSet(t.cfg)
	if err := rs.AddAllResources(); err != nil {
		return errors.Wrap(err, "couldn't add all resources to fargate resource set")
	}
	return t.stackManager.CreateStack(makeClusterStackName(t.cfg.Metadata.Name, t.cfg.Metadata.DisableStackPrefix), rs, nil, nil, errs)
}

// ensureFargateRoleStackExists creates fargate IAM resources if they
func ensureFargateRoleStackExists(
	cfg *api.ClusterConfig, provider api.ClusterProvider, stackManager manager.StackManager,
) error {
	if api.IsSetAndNonEmptyString(cfg.IAM.FargatePodExecutionRoleARN) {
		return nil
	}

	fargateStack, err := stackManager.GetFargateStack()
	if err != nil {
		return err
	}

	if fargateStack == nil {
		var taskTree tasks.TaskTree
		taskTree.Append(&createFargateStackTask{
			cfg:          cfg,
			provider:     provider,
			stackManager: stackManager,
		})

		errs := taskTree.DoAllSync()
		for _, e := range errs {
			logger.Critical("%s\n", e.Error())
		}

		if len(errs) > 0 {
			return errors.New("couldn't create fargate stack")
		}
	}
	return nil
}
