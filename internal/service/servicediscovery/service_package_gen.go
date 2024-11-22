// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package servicediscovery

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	servicediscovery_sdkv2 "github.com/aws/aws-sdk-go-v2/service/servicediscovery"
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
			Factory:  dataSourceDNSNamespace,
			TypeName: "aws_service_discovery_dns_namespace",
			Name:     "DNS Namespace",
		},
		{
			Factory:  dataSourceHTTPNamespace,
			TypeName: "aws_service_discovery_http_namespace",
			Name:     "HTTP Namespace",
		},
		{
			Factory:  dataSourceService,
			TypeName: "aws_service_discovery_service",
			Name:     "Service",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceHTTPNamespace,
			TypeName: "aws_service_discovery_http_namespace",
			Name:     "HTTP Namespace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceInstance,
			TypeName: "aws_service_discovery_instance",
			Name:     "Instance",
		},
		{
			Factory:  resourcePrivateDNSNamespace,
			TypeName: "aws_service_discovery_private_dns_namespace",
			Name:     "Private DNS Namespace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourcePublicDNSNamespace,
			TypeName: "aws_service_discovery_public_dns_namespace",
			Name:     "Public DNS Namespace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceService,
			TypeName: "aws_service_discovery_service",
			Name:     "Service",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ServiceDiscovery
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*servicediscovery_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return servicediscovery_sdkv2.NewFromConfig(cfg,
		servicediscovery_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
