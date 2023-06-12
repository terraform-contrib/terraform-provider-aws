// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package athena

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	athena_sdkv1 "github.com/aws/aws-sdk-go/service/athena"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct {
	config map[string]any
}

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
			Factory:  ResourceDataCatalog,
			TypeName: "aws_athena_data_catalog",
			Name:     "Data Catalog",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceDatabase,
			TypeName: "aws_athena_database",
		},
		{
			Factory:  ResourceNamedQuery,
			TypeName: "aws_athena_named_query",
		},
		{
			Factory:  ResourceWorkGroup,
			TypeName: "aws_athena_workgroup",
			Name:     "WorkGroup",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Athena
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context) (*athena_sdkv1.Athena, error) {
	sess := p.config["session"].(*session_sdkv1.Session)

	return athena_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
