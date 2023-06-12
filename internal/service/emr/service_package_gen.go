// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package emr

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	emr_sdkv1 "github.com/aws/aws-sdk-go/service/emr"
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
			Factory:  DataSourceReleaseLabels,
			TypeName: "aws_emr_release_labels",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceBlockPublicAccessConfiguration,
			TypeName: "aws_emr_block_public_access_configuration",
		},
		{
			Factory:  ResourceCluster,
			TypeName: "aws_emr_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceInstanceFleet,
			TypeName: "aws_emr_instance_fleet",
		},
		{
			Factory:  ResourceInstanceGroup,
			TypeName: "aws_emr_instance_group",
		},
		{
			Factory:  ResourceManagedScalingPolicy,
			TypeName: "aws_emr_managed_scaling_policy",
		},
		{
			Factory:  ResourceSecurityConfiguration,
			TypeName: "aws_emr_security_configuration",
		},
		{
			Factory:  ResourceStudio,
			TypeName: "aws_emr_studio",
			Name:     "Studio",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceStudioSessionMapping,
			TypeName: "aws_emr_studio_session_mapping",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.EMR
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context) (*emr_sdkv1.EMR, error) {
	sess := p.config["session"].(*session_sdkv1.Session)

	return emr_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
