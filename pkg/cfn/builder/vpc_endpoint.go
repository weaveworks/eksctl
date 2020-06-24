package builder

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	gfn "github.com/awslabs/goformation/cloudformation"
	"github.com/pkg/errors"
	"github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
)

type VPCEndpoint struct {
	provider           provider
	rs                 *resourceSet
	vpc                *gfn.Value
	additionalServices []string
	subnets            []SubnetResource
	clusterSG          *clusterSecurityGroup
}

type provider interface {
	EC2() ec2iface.EC2API
	Region() string
}

func NewVPCEndpointResourceSet(provider provider, rs *resourceSet, vpc *gfn.Value, subnets []SubnetResource, additionalServices []string, clusterSG *clusterSecurityGroup) *VPCEndpoint {
	return &VPCEndpoint{
		provider:           provider,
		vpc:                vpc,
		subnets:            subnets,
		additionalServices: additionalServices,
		rs:                 rs,
		clusterSG:          clusterSG,
	}
}

type VPCEndpointServiceDetails struct {
	ServiceName       string
	Service           string
	EndpointType      string
	AvailabilityZones []string
}

// AddResources adds resources for VPC endpoints
func (e *VPCEndpoint) AddResources() error {
	endpointServices := append(v1alpha5.DefaultEndpointServices(), e.additionalServices...)
	endpointServiceDetails, err := BuildVPCEndpointServices(e.provider.EC2(), e.provider.Region(), endpointServices)
	if err != nil {
		return errors.Wrap(err, "error building endpoint service details")
	}

	for _, endpointDetail := range endpointServiceDetails {
		endpoint := &gfn.AWSEC2VPCEndpoint{
			ServiceName:     gfn.NewString(endpointDetail.ServiceName),
			VpcId:           e.vpc,
			VpcEndpointType: gfn.NewString(endpointDetail.EndpointType),
		}

		if endpointDetail.EndpointType == ec2.VpcEndpointTypeGateway {
			endpoint.RouteTableIds = e.routeTableIDs()
		} else {
			endpoint.VpcEndpointType = gfn.NewString(ec2.VpcEndpointTypeInterface)
			endpoint.SubnetIds = e.subnetsForAZs(endpointDetail.AvailabilityZones)
			endpoint.PrivateDnsEnabled = gfn.NewBoolean(true)
			endpoint.SecurityGroupIds = []*gfn.Value{e.clusterSG.ClusterSharedNode}
		}

		resourceName := fmt.Sprintf("VPCEndpoint%s", strings.ToUpper(
			strings.ReplaceAll(endpointDetail.Service, ".", ""),
		))

		// TODO attach policy document
		e.rs.newResource(resourceName, endpoint)
	}
	return nil
}

func (e *VPCEndpoint) subnetsForAZs(azs []string) []*gfn.Value {
	var subnetRefs []*gfn.Value
	for _, az := range azs {
		for _, subnet := range e.subnets {
			if subnet.AvailabilityZone == az {
				subnetRefs = append(subnetRefs, subnet.Subnet)
			}
		}

	}
	return subnetRefs
}

func (e *VPCEndpoint) routeTableIDs() []*gfn.Value {
	var routeTableIDs []*gfn.Value
	for _, subnet := range e.subnets {
		routeTableIDs = append(routeTableIDs, subnet.RouteTable)
	}
	return routeTableIDs
}

func BuildVPCEndpointServices(ec2API ec2iface.EC2API, region string, endpoints []string) ([]VPCEndpointServiceDetails, error) {
	serviceNames := make([]string, len(endpoints))
	serviceRegionPrefix := fmt.Sprintf("com.amazonaws.%s.", region)
	for i, endpoint := range endpoints {
		serviceNames[i] = serviceRegionPrefix + endpoint
	}

	output, err := ec2API.DescribeVpcEndpointServices(&ec2.DescribeVpcEndpointServicesInput{
		ServiceNames: aws.StringSlice(serviceNames),
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("service-name"),
				Values: aws.StringSlice(serviceNames),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	serviceDetails := make([]VPCEndpointServiceDetails, len(output.ServiceDetails))

	for i, sd := range output.ServiceDetails {
		if len(sd.ServiceType) > 1 {
			return nil, errors.Errorf("endpoint service %q with multiple service types isn't supported", *sd.ServiceName)
		}

		serviceDetails[i] = VPCEndpointServiceDetails{
			ServiceName:       *sd.ServiceName,
			Service:           strings.TrimPrefix(*sd.ServiceName, serviceRegionPrefix),
			EndpointType:      *sd.ServiceType[0].ServiceType,
			AvailabilityZones: aws.StringValueSlice(sd.AvailabilityZones),
		}
	}

	return serviceDetails, nil
}