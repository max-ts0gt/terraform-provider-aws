// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package apigatewayv2

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
			Factory:  dataSourceAPI,
			TypeName: "aws_apigatewayv2_api",
			Name:     "API",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceAPIs,
			TypeName: "aws_apigatewayv2_apis",
			Name:     "APIs",
		},
		{
			Factory:  dataSourceExport,
			TypeName: "aws_apigatewayv2_export",
			Name:     "Export",
		},
		{
			Factory:  dataSourceVPCLink,
			TypeName: "aws_apigatewayv2_vpc_link",
			Name:     "VPC Link",
			Tags:     &types.ServicePackageResourceTags{},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAPI,
			TypeName: "aws_apigatewayv2_api",
			Name:     "API",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceAPIMapping,
			TypeName: "aws_apigatewayv2_api_mapping",
			Name:     "API Mapping",
		},
		{
			Factory:  resourceAuthorizer,
			TypeName: "aws_apigatewayv2_authorizer",
			Name:     "Authorizer",
		},
		{
			Factory:  resourceDeployment,
			TypeName: "aws_apigatewayv2_deployment",
			Name:     "Deployment",
		},
		{
			Factory:  resourceDomainName,
			TypeName: "aws_apigatewayv2_domain_name",
			Name:     "Domain Name",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceIntegration,
			TypeName: "aws_apigatewayv2_integration",
			Name:     "Integration",
		},
		{
			Factory:  resourceIntegrationResponse,
			TypeName: "aws_apigatewayv2_integration_response",
			Name:     "Integration Response",
		},
		{
			Factory:  resourceModel,
			TypeName: "aws_apigatewayv2_model",
			Name:     "Model",
		},
		{
			Factory:  resourceRoute,
			TypeName: "aws_apigatewayv2_route",
			Name:     "Route",
		},
		{
			Factory:  resourceRouteResponse,
			TypeName: "aws_apigatewayv2_route_response",
			Name:     "Route Response",
		},
		{
			Factory:  resourceStage,
			TypeName: "aws_apigatewayv2_stage",
			Name:     "Stage",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceVPCLink,
			TypeName: "aws_apigatewayv2_vpc_link",
			Name:     "VPC Link",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.APIGatewayV2
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
