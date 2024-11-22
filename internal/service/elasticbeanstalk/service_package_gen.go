// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	elasticbeanstalk_sdkv2 "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
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
			Factory:  dataSourceApplication,
			TypeName: "aws_elastic_beanstalk_application",
			Name:     "Application",
		},
		{
			Factory:  dataSourceHostedZone,
			TypeName: "aws_elastic_beanstalk_hosted_zone",
		},
		{
			Factory:  dataSourceSolutionStack,
			TypeName: "aws_elastic_beanstalk_solution_stack",
			Name:     "Solution Stack",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceApplication,
			TypeName: "aws_elastic_beanstalk_application",
			Name:     "Application",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceApplicationVersion,
			TypeName: "aws_elastic_beanstalk_application_version",
			Name:     "Application Version",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceConfigurationTemplate,
			TypeName: "aws_elastic_beanstalk_configuration_template",
			Name:     "Configuration Template",
		},
		{
			Factory:  resourceEnvironment,
			TypeName: "aws_elastic_beanstalk_environment",
			Name:     "Environment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ElasticBeanstalk
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*elasticbeanstalk_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return elasticbeanstalk_sdkv2.NewFromConfig(cfg,
		elasticbeanstalk_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
