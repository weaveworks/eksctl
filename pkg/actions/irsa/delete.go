package irsa

import (
	"fmt"

	"github.com/weaveworks/eksctl/pkg/utils/tasks"

	"github.com/kris-nova/logger"
	"github.com/weaveworks/eksctl/pkg/kubernetes"
)

func (m *Manager) Delete(serviceAccounts []string, plan, wait bool) error {
	taskTree, err := m.stackManager.NewTasksToDeleteIAMServiceAccounts(serviceAccounts, kubernetes.NewCachedClientSet(m.clientSet), wait)
	if err != nil {
		return err
	}
	taskTree.PlanMode = plan

	logger.Info(taskTree.Describe())
	if errs := taskTree.DoAllSync(); len(errs) > 0 {
		logger.Info("%d error(s) occurred and IAM Role stacks haven't been deleted properly, you may wish to check CloudFormation console", len(errs))
		for _, err := range errs {
			logger.Critical("%s\n", err.Error())
		}
		return fmt.Errorf("failed to delete iamserviceaccount(s)")
	}

	logPlanModeWarning(plan && taskTree.Len() > 0)
	return nil
}

func (m *Manager) DeleteTasks(serviceAccounts []string) (*tasks.TaskTree, error) {
	return m.stackManager.NewTasksToDeleteIAMServiceAccounts(serviceAccounts, kubernetes.NewCachedClientSet(m.clientSet), true)
}
