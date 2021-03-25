// +build integration

package apply

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/weaveworks/eksctl/pkg/utils/file"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/weaveworks/eksctl/integration/matchers"
	. "github.com/weaveworks/eksctl/integration/matchers"
	. "github.com/weaveworks/eksctl/integration/runner"
	"github.com/weaveworks/eksctl/integration/tests"
	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/eks"
	iamoidc "github.com/weaveworks/eksctl/pkg/iam/oidc"
	"github.com/weaveworks/eksctl/pkg/testutils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var params *tests.Params

func init() {
	// Call testing.Init() prior to tests.NewParams(), as otherwise -test.* will not be recognised. See also: https://golang.org/doc/go1.13#testing
	testing.Init()
	params = tests.NewParams("apply")
}

func TestCRUD(t *testing.T) {
	testutils.RegisterAndRun(t)
}

var _ = Describe("apply", func() {

	BeforeSuite(func() {
		params.KubeconfigTemp = false
		if params.KubeconfigPath == "" {
			wd, _ := os.Getwd()
			f, _ := ioutil.TempFile(wd, "kubeconfig-")
			params.KubeconfigPath = f.Name()
			params.KubeconfigTemp = true
		}

		if params.SkipCreate {
			params.ClusterName = "jk"
			fmt.Fprintf(GinkgoWriter, "will use existing cluster %s", params.ClusterName)
			if !file.Exists(params.KubeconfigPath) {
				// Generate the Kubernetes configuration that eksctl create
				// would have generated otherwise:
				cmd := params.EksctlUtilsCmd.WithArgs(
					"write-kubeconfig",
					"--verbose", "4",
					"--cluster", params.ClusterName,
					"--kubeconfig", params.KubeconfigPath,
				)
				Expect(cmd).To(RunSuccessfully())
			}
			return
		}

		fmt.Fprintf(GinkgoWriter, "Using kubeconfig: %s\n", params.KubeconfigPath)

		cmd := params.EksctlCreateCmd.WithArgs(
			"cluster",
			"--verbose", "4",
			"--name", params.ClusterName,
			"--tags", "alpha.eksctl.io/description=eksctl integration test",
			"--without-nodegroup",
			"--with-oidc",
			"--version", params.Version,
			"--kubeconfig", params.KubeconfigPath,
		)
		Expect(cmd).To(RunSuccessfully())
	})

	AfterSuite(func() {
		if !params.SkipCreate {
			params.DeleteClusters()
		}
		gexec.KillAndWait()
		if params.KubeconfigTemp {
			os.Remove(params.KubeconfigPath)
		}
		os.RemoveAll(params.TestDirectory)
	})

	Describe("IAMServiceAccounts", func() {
		Context("reconciling iamserviceaccounts", func() {
			var (
				cfg  *api.ClusterConfig
				ctl  *eks.ClusterProvider
				oidc *iamoidc.OpenIDConnectManager
				err  error
			)

			BeforeEach(func() {
				cfg = api.NewClusterConfig()
				cfg.Metadata = &api.ClusterMeta{
					Name:   params.ClusterName,
					Region: params.Region,
				}
				cfg.IAM.WithOIDC = aws.Bool(true)

				ctl, err = eks.New(&api.ProviderConfig{Region: params.Region}, cfg)
				Expect(err).NotTo(HaveOccurred())
				err = ctl.RefreshClusterStatus(cfg)
				Expect(err).ShouldNot(HaveOccurred())
				oidc, err = ctl.NewOpenIDConnectManager(cfg)
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("should reconcile the cluster state to match the IAMServiceAccounts listed in the cluster config", func() {
				exists, err := oidc.CheckProviderExists()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(exists).To(BeTrue())

				cfg.IAM.ServiceAccounts = []*api.ClusterIAMServiceAccount{
					{
						ClusterIAMMeta: api.ClusterIAMMeta{
							Namespace: "default",
							Name:      "create",
						},
						AttachPolicyARNs: []string{"arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"},
					},
					{
						ClusterIAMMeta: api.ClusterIAMMeta{
							Namespace: "default",
							Name:      "no-sa",
						},
						AttachPolicyARNs: []string{"arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"},
					},
					{
						ClusterIAMMeta: api.ClusterIAMMeta{
							Namespace: "default",
							Name:      "update",
						},
						AttachPolicyARNs: []string{"arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"},
					},
					{
						ClusterIAMMeta: api.ClusterIAMMeta{
							Namespace: "default",
							Name:      "delete",
						},
						AttachPolicyARNs: []string{"arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"},
					},
				}

				configFilePath := writeConfig(cfg)

				cmd := params.EksctlApplyCmd.WithArgs(
					"--config-file", configFilePath,
					"--approve",
				)
				Expect(cmd).To(RunSuccessfully())
				awsSession := NewSession(params.Region)

				stackNamePrefix := fmt.Sprintf("eksctl-%s-addon-iamserviceaccount-", params.ClusterName)

				Expect(awsSession).To(HaveExistingStack(stackNamePrefix + "default-create"))
				Expect(awsSession).To(HaveExistingStack(stackNamePrefix + "default-update"))
				Expect(awsSession).To(HaveExistingStack(stackNamePrefix + "default-delete"))
				Expect(awsSession).To(HaveExistingStack(stackNamePrefix + "default-no-sa"))

				clientSet, err := ctl.NewStdClientSet(cfg)
				Expect(err).ShouldNot(HaveOccurred())

				for _, saName := range []string{"create", "update", "delete", "no-sa"} {
					sa, err := clientSet.CoreV1().ServiceAccounts(metav1.NamespaceDefault).Get(context.TODO(), saName, metav1.GetOptions{})
					Expect(err).ShouldNot(HaveOccurred())

					Expect(sa.Annotations).To(HaveLen(1))
					Expect(sa.Annotations).To(HaveKey(api.AnnotationEKSRoleARN))
					Expect(sa.Annotations[api.AnnotationEKSRoleARN]).To(MatchRegexp("^arn:aws:iam::.*:role/eksctl-" + truncate(params.ClusterName) + ".*$"))
				}

				By("removing one IAMServiceAccounts,creating a new one, updating the policys of one and removing one of the kubernetes SA")

				err = clientSet.CoreV1().ServiceAccounts("default").Delete(context.Background(), "no-sa", metav1.DeleteOptions{})
				Expect(err).NotTo(HaveOccurred())

				cfg.IAM.ServiceAccounts = []*api.ClusterIAMServiceAccount{
					{
						ClusterIAMMeta: api.ClusterIAMMeta{
							Namespace: "default",
							Name:      "create",
						},
						AttachPolicyARNs: []string{"arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"},
					},
					{
						ClusterIAMMeta: api.ClusterIAMMeta{
							Namespace: "default",
							Name:      "new",
						},
						AttachPolicyARNs: []string{"arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"},
					},
					{
						ClusterIAMMeta: api.ClusterIAMMeta{
							Namespace: "default",
							Name:      "no-sa",
						},
						AttachPolicyARNs: []string{"arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"},
					},
					{
						ClusterIAMMeta: api.ClusterIAMMeta{
							Namespace: "default",
							Name:      "update",
						},
						AttachPolicyARNs: []string{
							"arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess",
							"arn:aws:iam::aws:policy/AmazonElastiCacheFullAccess",
						},
					},
				}

				configFilePath = writeConfig(cfg)

				cmd = params.EksctlApplyCmd.WithArgs(
					"--config-file", configFilePath,
					"--approve",
				)
				Expect(cmd).To(RunSuccessfully())
				awsSession = NewSession(params.Region)

				stackNamePrefix = fmt.Sprintf("eksctl-%s-addon-iamserviceaccount-", params.ClusterName)

				Expect(awsSession).To(HaveExistingStack(stackNamePrefix + "default-create"))
				Expect(awsSession).To(HaveExistingStack(stackNamePrefix + "default-new"))
				Expect(awsSession).To(HaveExistingStack(stackNamePrefix + "default-update"))
				Expect(awsSession).To(HaveExistingStack(stackNamePrefix + "default-no-sa"))
				Eventually(func() bool {
					stackName := stackNamePrefix + "default-delete"
					exists, err := matchers.StackExists(stackName, awsSession)
					Expect(err).NotTo(HaveOccurred())
					return exists
				}, time.Minute, time.Second*10).Should(BeFalse())

				for _, saName := range []string{"create", "update", "new", "no-sa"} {
					sa, err := clientSet.CoreV1().ServiceAccounts(metav1.NamespaceDefault).Get(context.TODO(), saName, metav1.GetOptions{})
					Expect(err).ShouldNot(HaveOccurred())

					Expect(sa.Annotations).To(HaveLen(1))
					Expect(sa.Annotations).To(HaveKey(api.AnnotationEKSRoleARN))
					Expect(sa.Annotations[api.AnnotationEKSRoleARN]).To(MatchRegexp("^arn:aws:iam::.*:role/eksctl-" + truncate(params.ClusterName) + ".*$"))
				}

				_, err = clientSet.CoreV1().ServiceAccounts(metav1.NamespaceDefault).Get(context.TODO(), "delete", metav1.GetOptions{})
				Expect(err).Should(HaveOccurred())
				Expect(apierrors.IsNotFound(err)).To(BeTrue())

			})
		})
	})
})

func writeConfig(cfg *api.ClusterConfig) string {
	configData, err := json.Marshal(&cfg)
	Expect(err).NotTo(HaveOccurred())
	configFile, err := ioutil.TempFile("", "")
	Expect(err).NotTo(HaveOccurred())
	Expect(ioutil.WriteFile(configFile.Name(), configData, 0755)).To(Succeed())
	return configFile.Name()
}

func truncate(clusterName string) string {
	// CloudFormation seems to truncate long cluster names at 37 characters:
	if len(clusterName) > 37 {
		return clusterName[:37]
	}
	return clusterName
}
