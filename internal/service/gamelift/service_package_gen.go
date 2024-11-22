// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package gamelift

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	gamelift_sdkv2 "github.com/aws/aws-sdk-go-v2/service/gamelift"
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
			Factory:  resourceAlias,
			TypeName: "aws_gamelift_alias",
			Name:     "Alias",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceBuild,
			TypeName: "aws_gamelift_build",
			Name:     "Build",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceFleet,
			TypeName: "aws_gamelift_fleet",
			Name:     "Fleet",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceGameServerGroup,
			TypeName: "aws_gamelift_game_server_group",
			Name:     "Game Server Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceGameSessionQueue,
			TypeName: "aws_gamelift_game_session_queue",
			Name:     "Game Session Queue",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceScript,
			TypeName: "aws_gamelift_script",
			Name:     "Script",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.GameLift
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*gamelift_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return gamelift_sdkv2.NewFromConfig(cfg,
		gamelift_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
