// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package mq

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	mq_sdkv2 "github.com/aws/aws-sdk-go-v2/service/mq"
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
			Factory:  dataSourceBroker,
			TypeName: "aws_mq_broker",
			Name:     "Broker",
		},
		{
			Factory:  dataSourceBrokerEngineTypes,
			TypeName: "aws_mq_broker_engine_types",
			Name:     "Broker Engine Types",
		},
		{
			Factory:  dataSourceBrokerInstanceTypeOfferings,
			TypeName: "aws_mq_broker_instance_type_offerings",
			Name:     "Broker Instance Type Offerings",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceBroker,
			TypeName: "aws_mq_broker",
			Name:     "Broker",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceConfiguration,
			TypeName: "aws_mq_configuration",
			Name:     "Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.MQ
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*mq_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return mq_sdkv2.NewFromConfig(cfg,
		mq_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
