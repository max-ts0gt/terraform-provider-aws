// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package cloudformation

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
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceExport,
			TypeName: "aws_cloudformation_export",
			Name:     "Export",
		},
		{
			Factory:  dataSourceStack,
			TypeName: "aws_cloudformation_stack",
			Name:     "Stack",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceType,
			TypeName: "aws_cloudformation_type",
			Name:     "Type",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceStack,
			TypeName: "aws_cloudformation_stack",
			Name:     "Stack",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  resourceStackSet,
			TypeName: "aws_cloudformation_stack_set",
			Name:     "Stack Set",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  resourceStackSetInstance,
			TypeName: "aws_cloudformation_stack_set_instance",
			Name:     "Stack Set Instance",
		},
		{
			Factory:  resourceType,
			TypeName: "aws_cloudformation_type",
			Name:     "Type",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CloudFormation
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
