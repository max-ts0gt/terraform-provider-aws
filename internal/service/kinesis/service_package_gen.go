// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package kinesis

import (
	"context"

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
			Factory: newResourcePolicyResource,
			Name:    "Resource Policy",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceStream,
			TypeName: "aws_kinesis_stream",
		},
		{
			Factory:  dataSourceStreamConsumer,
			TypeName: "aws_kinesis_stream_consumer",
			Name:     "Stream Consumer",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceStream,
			TypeName: "aws_kinesis_stream",
			Name:     "Stream",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrName,
			},
		},
		{
			Factory:  resourceStreamConsumer,
			TypeName: "aws_kinesis_stream_consumer",
			Name:     "Stream Consumer",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Kinesis
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
