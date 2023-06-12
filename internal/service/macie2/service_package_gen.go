// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package macie2

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	macie2_sdkv1 "github.com/aws/aws-sdk-go/service/macie2"
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
			Factory:  ResourceAccount,
			TypeName: "aws_macie2_account",
		},
		{
			Factory:  ResourceClassificationExportConfiguration,
			TypeName: "aws_macie2_classification_export_configuration",
		},
		{
			Factory:  ResourceClassificationJob,
			TypeName: "aws_macie2_classification_job",
			Name:     "Classification Job",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceCustomDataIdentifier,
			TypeName: "aws_macie2_custom_data_identifier",
			Name:     "Custom Data Identifier",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceFindingsFilter,
			TypeName: "aws_macie2_findings_filter",
			Name:     "Findings Filter",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceInvitationAccepter,
			TypeName: "aws_macie2_invitation_accepter",
		},
		{
			Factory:  ResourceMember,
			TypeName: "aws_macie2_member",
			Name:     "Member",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceOrganizationAdminAccount,
			TypeName: "aws_macie2_organization_admin_account",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Macie2
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context) (*macie2_sdkv1.Macie2, error) {
	sess := p.config["session"].(*session_sdkv1.Session)

	return macie2_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
