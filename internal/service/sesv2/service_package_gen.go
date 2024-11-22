// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package sesv2

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	sesv2_sdkv2 "github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/max-ts0gt/terraform-provider-aws/internal/conns"
	"github.com/max-ts0gt/terraform-provider-aws/internal/types"
	"github.com/max-ts0gt/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceConfigurationSet,
			TypeName: "aws_sesv2_configuration_set",
		},
		{
			Factory:  DataSourceDedicatedIPPool,
			TypeName: "aws_sesv2_dedicated_ip_pool",
		},
		{
			Factory:  DataSourceEmailIdentity,
			TypeName: "aws_sesv2_email_identity",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  DataSourceEmailIdentityMailFromAttributes,
			TypeName: "aws_sesv2_email_identity_mail_from_attributes",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAccountVDMAttributes,
			TypeName: "aws_sesv2_account_vdm_attributes",
			Name:     "Account VDM Attributes",
		},
		{
			Factory:  ResourceConfigurationSet,
			TypeName: "aws_sesv2_configuration_set",
			Name:     "Configuration Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceConfigurationSetEventDestination,
			TypeName: "aws_sesv2_configuration_set_event_destination",
		},
		{
			Factory:  ResourceContactList,
			TypeName: "aws_sesv2_contact_list",
			Name:     "Contact List",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceDedicatedIPAssignment,
			TypeName: "aws_sesv2_dedicated_ip_assignment",
		},
		{
			Factory:  ResourceDedicatedIPPool,
			TypeName: "aws_sesv2_dedicated_ip_pool",
			Name:     "Dedicated IP Pool",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceEmailIdentity,
			TypeName: "aws_sesv2_email_identity",
			Name:     "Email Identity",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceEmailIdentityFeedbackAttributes,
			TypeName: "aws_sesv2_email_identity_feedback_attributes",
		},
		{
			Factory:  ResourceEmailIdentityMailFromAttributes,
			TypeName: "aws_sesv2_email_identity_mail_from_attributes",
		},
		{
			Factory:  ResourceEmailIdentityPolicy,
			TypeName: "aws_sesv2_email_identity_policy",
			Name:     "Email Identity Policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SESV2
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*sesv2_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return sesv2_sdkv2.NewFromConfig(cfg,
		sesv2_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
