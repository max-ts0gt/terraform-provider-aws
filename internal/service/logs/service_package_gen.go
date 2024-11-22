// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package logs

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	cloudwatchlogs_sdkv2 "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
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
			Factory:  dataSourceDataProtectionPolicyDocument,
			TypeName: "aws_cloudwatch_log_data_protection_policy_document",
		},
		{
			Factory:  dataSourceGroup,
			TypeName: "aws_cloudwatch_log_group",
			Name:     "Log Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceGroups,
			TypeName: "aws_cloudwatch_log_groups",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccountPolicy,
			TypeName: "aws_cloudwatch_log_account_policy",
			Name:     "Account Policy",
		},
		{
			Factory:  resourceDataProtectionPolicy,
			TypeName: "aws_cloudwatch_log_data_protection_policy",
		},
		{
			Factory:  resourceDestination,
			TypeName: "aws_cloudwatch_log_destination",
			Name:     "Destination",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceDestinationPolicy,
			TypeName: "aws_cloudwatch_log_destination_policy",
		},
		{
			Factory:  resourceGroup,
			TypeName: "aws_cloudwatch_log_group",
			Name:     "Log Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceMetricFilter,
			TypeName: "aws_cloudwatch_log_metric_filter",
		},
		{
			Factory:  resourceResourcePolicy,
			TypeName: "aws_cloudwatch_log_resource_policy",
		},
		{
			Factory:  resourceStream,
			TypeName: "aws_cloudwatch_log_stream",
		},
		{
			Factory:  resourceSubscriptionFilter,
			TypeName: "aws_cloudwatch_log_subscription_filter",
		},
		{
			Factory:  resourceQueryDefinition,
			TypeName: "aws_cloudwatch_query_definition",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Logs
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*cloudwatchlogs_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return cloudwatchlogs_sdkv2.NewFromConfig(cfg,
		cloudwatchlogs_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
