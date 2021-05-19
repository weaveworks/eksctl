// +build integration

package unowned_clusters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/weaveworks/eksctl/pkg/eks"

	cfn "github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/aws/aws-sdk-go/aws"

	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"

	awseks "github.com/aws/aws-sdk-go/service/eks"
	. "github.com/weaveworks/eksctl/integration/runner"
	"github.com/weaveworks/eksctl/integration/tests"
	"github.com/weaveworks/eksctl/pkg/testutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var params *tests.Params

func init() {
	// Call testing.Init() prior to tests.NewParams(), as otherwise -test.* will not be recognised. See also: https://golang.org/doc/go1.13#testing
	testing.Init()
	// "unowned_clusters" lead to names longer than allowed for CF stacks
	params = tests.NewParams("uc")
}

func TestE2E(t *testing.T) {
	testutils.RegisterAndRun(t)
}

var _ = Describe("(Integration) [non-eksctl cluster & nodegroup support]", func() {
	var (
		stackName, version, ng1, mng1 string
		ctl                           api.ClusterProvider
		configFile                    *os.File
		cfg                           *api.ClusterConfig
		kmsKeyARN                     *string
	)

	BeforeSuite(func() {
		ng1 = "ng-1"
		mng1 = "mng-1"
		version = "1.19"
		stackName = fmt.Sprintf("eksctl-%s", params.ClusterName)
		params.ClusterName = "JKJKJK" + params.ClusterName
		cfg = &api.ClusterConfig{
			TypeMeta: api.ClusterConfigTypeMeta(),
			Metadata: &api.ClusterMeta{
				Version: version,
				Name:    params.ClusterName,
				Region:  params.Region,
			},
		}

		var err error
		configFile, err = ioutil.TempFile("", "")
		Expect(err).NotTo(HaveOccurred())

		if !params.SkipCreate {
			clusterProvider, err := eks.New(&api.ProviderConfig{Region: params.Region}, cfg)
			Expect(err).NotTo(HaveOccurred())
			ctl = clusterProvider.Provider
			cfg.VPC = createClusterWithNodeGroup(params.ClusterName, stackName, mng1, version, ctl)

			kmsClient := kms.New(ctl.ConfigProvider())
			output, err := kmsClient.CreateKey(&kms.CreateKeyInput{
				Description: aws.String(fmt.Sprintf("Key to test KMS encryption on EKS cluster %s", params.ClusterName)),
			})
			Expect(err).NotTo(HaveOccurred())
			kmsKeyARN = output.KeyMetadata.Arn
		}
	})

	AfterSuite(func() {
		if !params.SkipCreate && !params.SkipDelete {
			deleteStack(stackName, ctl)

			kmsClient := kms.New(ctl.ConfigProvider())
			_, err := kmsClient.ScheduleKeyDeletion(&kms.ScheduleKeyDeletionInput{
				KeyId:               kmsKeyARN,
				PendingWindowInDays: aws.Int64(7),
			})
			Expect(err).NotTo(HaveOccurred())
		}
		Expect(os.RemoveAll(configFile.Name())).To(Succeed())

	})

	It("supports creating nodegroups", func() {
		cfg.NodeGroups = []*api.NodeGroup{{
			NodeGroupBase: &api.NodeGroupBase{
				Name: ng1,
			}},
		}
		// write config file so that the nodegroup creates have access to the vpc spec
		configData, err := json.Marshal(&cfg)
		Expect(err).NotTo(HaveOccurred())
		Expect(ioutil.WriteFile(configFile.Name(), configData, 0755)).To(Succeed())
		cmd := params.EksctlCreateNodegroupCmd.
			WithArgs(
				"--config-file", configFile.Name(),
				"--verbose", "2",
			)
		Expect(cmd).To(RunSuccessfully())
	})

	It("supports deleting clusters", func() {
		if params.SkipDelete {
			Skip("params.SkipDelete is true")
		}
		By("deleting the cluster")
		cmd := params.EksctlDeleteCmd.
			WithArgs(
				"cluster",
				"--name", params.ClusterName,
				"--timeout", "1h",
				"--verbose", "3",
			)
		Expect(cmd).To(RunSuccessfully())
	})
})

func createClusterWithNodeGroup(clusterName, stackName, ng1, version string, ctl api.ClusterProvider) *api.ClusterVPC {
	timeoutDuration := time.Minute * 30
	publicSubnets, privateSubnets, clusterRoleArn, nodeRoleArn, vpcID, securityGroup := createVPCAndRole(stackName, ctl)

	_, err := ctl.EKS().CreateCluster(&awseks.CreateClusterInput{
		Name: &clusterName,
		ResourcesVpcConfig: &awseks.VpcConfigRequest{
			SubnetIds: aws.StringSlice(append(publicSubnets, privateSubnets...)),
		},
		RoleArn: &clusterRoleArn,
		Version: aws.String(version),
	})
	Expect(err).NotTo(HaveOccurred())
	Eventually(func() string {
		out, err := ctl.EKS().DescribeCluster(&awseks.DescribeClusterInput{
			Name: &clusterName,
		})
		Expect(err).NotTo(HaveOccurred())
		return *out.Cluster.Status
	}, timeoutDuration, time.Second*30).Should(Equal("ACTIVE"))

	newVPC := api.NewClusterVPC()
	newVPC.ID = vpcID
	newVPC.SecurityGroup = securityGroup
	newVPC.Subnets = &api.ClusterSubnets{
		Public: api.AZSubnetMapping{
			"public1": api.AZSubnetSpec{
				ID: publicSubnets[0],
			},
			"public2": api.AZSubnetSpec{
				ID: publicSubnets[1],
			},
			"public3": api.AZSubnetSpec{
				ID: publicSubnets[2],
			},
		},
		Private: api.AZSubnetMapping{
			"private4": api.AZSubnetSpec{
				ID: privateSubnets[0],
			},
			"private5": api.AZSubnetSpec{
				ID: privateSubnets[1],
			},
			"private6": api.AZSubnetSpec{
				ID: privateSubnets[2],
			},
		},
	}

	_, err = ctl.EKS().CreateNodegroup(&awseks.CreateNodegroupInput{
		NodegroupName: &ng1,
		ClusterName:   &clusterName,
		NodeRole:      &nodeRoleArn,
		Subnets:       aws.StringSlice(publicSubnets),
		ScalingConfig: &awseks.NodegroupScalingConfig{
			MaxSize:     aws.Int64(1),
			DesiredSize: aws.Int64(1),
			MinSize:     aws.Int64(1),
		},
	})
	Expect(err).NotTo(HaveOccurred())
	Eventually(func() string {
		out, err := ctl.EKS().DescribeNodegroup(&awseks.DescribeNodegroupInput{
			ClusterName:   &clusterName,
			NodegroupName: &ng1,
		})
		Expect(err).NotTo(HaveOccurred())
		return *out.Nodegroup.Status
	}, timeoutDuration, time.Second*30).Should(Equal("ACTIVE"))

	return newVPC
}

func createVPCAndRole(stackName string, ctl api.ClusterProvider) ([]string, []string, string, string, string, string) {
	templateBody, err := ioutil.ReadFile("cf-template.yaml")
	Expect(err).NotTo(HaveOccurred())
	createStackInput := &cfn.CreateStackInput{
		StackName: &stackName,
	}
	createStackInput.SetTemplateBody(string(templateBody))
	createStackInput.SetCapabilities(aws.StringSlice([]string{cfn.CapabilityCapabilityIam}))
	createStackInput.SetCapabilities(aws.StringSlice([]string{cfn.CapabilityCapabilityNamedIam}))

	_, err = ctl.CloudFormation().CreateStack(createStackInput)
	Expect(err).NotTo(HaveOccurred())

	var describeStackOut *cfn.DescribeStacksOutput
	Eventually(func() string {
		describeStackOut, err = ctl.CloudFormation().DescribeStacks(&cfn.DescribeStacksInput{
			StackName: &stackName,
		})
		Expect(err).NotTo(HaveOccurred())
		return *describeStackOut.Stacks[0].StackStatus
	}, time.Minute*10, time.Second*15).Should(Equal(cfn.StackStatusCreateComplete))

	var clusterRoleARN, nodeRoleARN, vpcID string
	var publicSubnets, privateSubnets, securityGroups []string
	for _, output := range describeStackOut.Stacks[0].Outputs {
		switch *output.OutputKey {
		case "ClusterRoleARN":
			clusterRoleARN = *output.OutputValue
		case "NodeRoleARN":
			nodeRoleARN = *output.OutputValue
		case "PublicSubnetIds":
			publicSubnets = strings.Split(*output.OutputValue, ",")
		case "PrivateSubnetIds":
			privateSubnets = strings.Split(*output.OutputValue, ",")
		case "VpcId":
			vpcID = *output.OutputValue
		case "SecurityGroups":
			securityGroups = strings.Split(*output.OutputValue, ",")
		}
	}

	return publicSubnets, privateSubnets, clusterRoleARN, nodeRoleARN, vpcID, securityGroups[0]
}

func deleteStack(stackName string, ctl api.ClusterProvider) {
	deleteStackInput := &cfn.DeleteStackInput{
		StackName: &stackName,
	}

	_, err := ctl.CloudFormation().DeleteStack(deleteStackInput)
	Expect(err).NotTo(HaveOccurred())
}
