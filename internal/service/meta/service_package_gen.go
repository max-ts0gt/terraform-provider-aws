// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package meta

import (
	"context"

	"github.com/max-ts0gt/terraform-provider-aws/internal/conns"
	"github.com/max-ts0gt/terraform-provider-aws/internal/types"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceARN,
		},
		{
			Factory: newDataSourceBillingServiceAccount,
		},
		{
			Factory: newDataSourceDefaultTags,
		},
		{
			Factory: newDataSourceIPRanges,
		},
		{
			Factory: newDataSourcePartition,
		},
		{
			Factory: newDataSourceRegion,
		},
		{
			Factory: newDataSourceRegions,
			Name:    "Regions",
		},
		{
			Factory: newDataSourceService,
		},
		{
			Factory: newServicePrincipalDataSource,
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return "meta"
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
