// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ce

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	costexplorer_sdkv1 "github.com/aws/aws-sdk-go/service/costexplorer"
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
			Factory:  DataSourceCostCategory,
			TypeName: "aws_ce_cost_category",
		},
		{
			Factory:  DataSourceTags,
			TypeName: "aws_ce_tags",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAnomalyMonitor,
			TypeName: "aws_ce_anomaly_monitor",
			Name:     "Anomaly Monitor",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceAnomalySubscription,
			TypeName: "aws_ce_anomaly_subscription",
			Name:     "Anomaly Subscription",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceCostAllocationTag,
			TypeName: "aws_ce_cost_allocation_tag",
		},
		{
			Factory:  ResourceCostCategory,
			TypeName: "aws_ce_cost_category",
			Name:     "Cost Category",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CE
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context) (*costexplorer_sdkv1.CostExplorer, error) {
	sess := p.config["session"].(*session_sdkv1.Session)

	return costexplorer_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
