// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package securityhub

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	securityhub_sdkv2 "github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/max-ts0gt/terraform-provider-aws/internal/conns"
	"github.com/max-ts0gt/terraform-provider-aws/internal/types"
	"github.com/max-ts0gt/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newAutomationRuleResource,
			Name:    "Automation Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccount,
			TypeName: "aws_securityhub_account",
			Name:     "Account",
		},
		{
			Factory:  resourceActionTarget,
			TypeName: "aws_securityhub_action_target",
			Name:     "Action Target",
		},
		{
			Factory:  resourceConfigurationPolicy,
			TypeName: "aws_securityhub_configuration_policy",
			Name:     "Configuration Policy",
		},
		{
			Factory:  resourceConfigurationPolicyAssociation,
			TypeName: "aws_securityhub_configuration_policy_association",
			Name:     "Configuration Policy Association",
		},
		{
			Factory:  resourceFindingAggregator,
			TypeName: "aws_securityhub_finding_aggregator",
			Name:     "Finding Aggregator",
		},
		{
			Factory:  resourceInsight,
			TypeName: "aws_securityhub_insight",
			Name:     "Insight",
		},
		{
			Factory:  resourceInviteAccepter,
			TypeName: "aws_securityhub_invite_accepter",
			Name:     "Invite Accepter",
		},
		{
			Factory:  resourceMember,
			TypeName: "aws_securityhub_member",
			Name:     "Member",
		},
		{
			Factory:  resourceOrganizationAdminAccount,
			TypeName: "aws_securityhub_organization_admin_account",
			Name:     "Organization Admin Account",
		},
		{
			Factory:  resourceOrganizationConfiguration,
			TypeName: "aws_securityhub_organization_configuration",
			Name:     "Organization Configuration",
		},
		{
			Factory:  resourceProductSubscription,
			TypeName: "aws_securityhub_product_subscription",
			Name:     "Product Subscription",
		},
		{
			Factory:  resourceStandardsControl,
			TypeName: "aws_securityhub_standards_control",
			Name:     "Standards Control",
		},
		{
			Factory:  resourceStandardsSubscription,
			TypeName: "aws_securityhub_standards_subscription",
			Name:     "Standards Subscription",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SecurityHub
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*securityhub_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return securityhub_sdkv2.NewFromConfig(cfg,
		securityhub_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
