// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package secretsmanager

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	secretsmanager_sdkv1 "github.com/aws/aws-sdk-go/service/secretsmanager"
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
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceRandomPassword,
			TypeName: "aws_secretsmanager_random_password",
		},
		{
			Factory:  DataSourceSecret,
			TypeName: "aws_secretsmanager_secret",
		},
		{
			Factory:  DataSourceSecretRotation,
			TypeName: "aws_secretsmanager_secret_rotation",
		},
		{
			Factory:  DataSourceSecretVersion,
			TypeName: "aws_secretsmanager_secret_version",
		},
		{
			Factory:  DataSourceSecrets,
			TypeName: "aws_secretsmanager_secrets",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceSecret,
			TypeName: "aws_secretsmanager_secret",
			Name:     "Secret",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceSecretPolicy,
			TypeName: "aws_secretsmanager_secret_policy",
		},
		{
			Factory:  ResourceSecretRotation,
			TypeName: "aws_secretsmanager_secret_rotation",
		},
		{
			Factory:  ResourceSecretVersion,
			TypeName: "aws_secretsmanager_secret_version",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SecretsManager
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context) (*secretsmanager_sdkv1.SecretsManager, error) {
	sess := p.config["session"].(*session_sdkv1.Session)

	return secretsmanager_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
