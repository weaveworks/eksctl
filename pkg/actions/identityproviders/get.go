package identityproviders

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/pkg/errors"

	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
)

// Summary holds the known info about this provider
type Summary struct {
	Type      api.IdentityProviderType
	Name      string
	ClientID  string
	IssuerURL string
	Status    string
	Arn       string
}

type GetIdentityProvidersOptions struct {
	Name string
}

func (ipm *IdentityProviderManager) Get(options GetIdentityProvidersOptions) ([]Summary, error) {
	summaries := []Summary{}
	var configs []*eks.IdentityProviderConfig

	input := eks.ListIdentityProviderConfigsInput{
		ClusterName: aws.String(ipm.metadata.Name),
	}
	list, err := ipm.eksAPI.ListIdentityProviderConfigs(&input)
	if err != nil {
		return summaries, err
	}

	if options.Name == "" {
		configs = list.IdentityProviderConfigs
	} else {
		var getCfg *eks.IdentityProviderConfig
		for _, cfg := range list.IdentityProviderConfigs {
			if aws.StringValue(cfg.Name) == options.Name {
				getCfg = &eks.IdentityProviderConfig{
					Name: aws.String(options.Name),
					Type: cfg.Type,
				}
			}
		}
		if getCfg == nil {
			return summaries, errors.Errorf("couldn't find identity provider %s", options.Name)
		}
		configs = []*eks.IdentityProviderConfig{getCfg}
	}
	for _, idp := range configs {
		input := eks.DescribeIdentityProviderConfigInput{
			ClusterName:            aws.String(ipm.metadata.Name),
			IdentityProviderConfig: idp,
		}
		idP, err := ipm.eksAPI.DescribeIdentityProviderConfig(&input)
		if err != nil {
			return summaries, err
		}
		if cfg := idP.IdentityProviderConfig.Oidc; cfg != nil {
			summaries = append(summaries, Summary{
				Type:      api.OIDCIdentityProviderType,
				Name:      aws.StringValue(cfg.IdentityProviderConfigName),
				ClientID:  aws.StringValue(cfg.ClientId),
				IssuerURL: aws.StringValue(cfg.IssuerUrl),
				Status:    aws.StringValue(cfg.Status),
				Arn:       aws.StringValue(cfg.IdentityProviderConfigArn),
			})
		}
	}
	return summaries, nil
}
