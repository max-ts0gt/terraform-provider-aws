// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package lexv2models

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	lexmodelsv2_sdkv2 "github.com/aws/aws-sdk-go-v2/service/lexmodelsv2"
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
			Factory: newResourceBot,
			Name:    "Bot",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory: newResourceBotLocale,
			Name:    "Bot Locale",
		},
		{
			Factory: newResourceBotVersion,
			Name:    "Bot Version",
		},
		{
			Factory: newResourceIntent,
			Name:    "Intent",
		},
		{
			Factory: newResourceSlot,
			Name:    "Slot",
		},
		{
			Factory: newResourceSlotType,
			Name:    "Slot Type",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.LexV2Models
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*lexmodelsv2_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return lexmodelsv2_sdkv2.NewFromConfig(cfg,
		lexmodelsv2_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
