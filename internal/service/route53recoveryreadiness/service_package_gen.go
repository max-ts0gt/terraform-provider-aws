// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package route53recoveryreadiness

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
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceCell,
			TypeName: "aws_route53recoveryreadiness_cell",
			Name:     "Cell",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceReadinessCheck,
			TypeName: "aws_route53recoveryreadiness_readiness_check",
			Name:     "Readiness Check",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceRecoveryGroup,
			TypeName: "aws_route53recoveryreadiness_recovery_group",
			Name:     "Recovery Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceResourceSet,
			TypeName: "aws_route53recoveryreadiness_resource_set",
			Name:     "Resource Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Route53RecoveryReadiness
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
