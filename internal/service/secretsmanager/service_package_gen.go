// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package secretsmanager

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	secretsmanager_sdkv2 "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/max-ts0gt/terraform-provider-aws/internal/conns"
	"github.com/max-ts0gt/terraform-provider-aws/internal/types"
	"github.com/max-ts0gt/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceSecretVersions,
			Name:    "Secret Versions",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceRandomPassword,
			TypeName: "aws_secretsmanager_random_password",
			Name:     "Random Password",
		},
		{
			Factory:  dataSourceSecret,
			TypeName: "aws_secretsmanager_secret",
			Name:     "Secret",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceSecretRotation,
			TypeName: "aws_secretsmanager_secret_rotation",
			Name:     "Secret Rotation",
		},
		{
			Factory:  dataSourceSecretVersion,
			TypeName: "aws_secretsmanager_secret_version",
			Name:     "Secret Version",
		},
		{
			Factory:  dataSourceSecrets,
			TypeName: "aws_secretsmanager_secrets",
			Name:     "Secrets",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceSecret,
			TypeName: "aws_secretsmanager_secret",
			Name:     "Secret",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceSecretPolicy,
			TypeName: "aws_secretsmanager_secret_policy",
			Name:     "Secret Policy",
		},
		{
			Factory:  resourceSecretRotation,
			TypeName: "aws_secretsmanager_secret_rotation",
			Name:     "Secret Rotation",
		},
		{
			Factory:  resourceSecretVersion,
			TypeName: "aws_secretsmanager_secret_version",
			Name:     "Secret Version",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SecretsManager
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*secretsmanager_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return secretsmanager_sdkv2.NewFromConfig(cfg,
		secretsmanager_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
