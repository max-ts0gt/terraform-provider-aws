// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package s3

import (
	"context"

	"github.com/max-ts0gt/terraform-provider-aws/internal/conns"
	"github.com/max-ts0gt/terraform-provider-aws/internal/types"
	"github.com/max-ts0gt/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDirectoryBucketsDataSource,
			Name:    "Directory Buckets",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newDirectoryBucketResource,
			Name:    "Directory Bucket",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCanonicalUserID,
			TypeName: "aws_canonical_user_id",
			Name:     "Canonical User ID",
		},
		{
			Factory:  dataSourceBucket,
			TypeName: "aws_s3_bucket",
			Name:     "Bucket",
		},
		{
			Factory:  dataSourceBucketObject,
			TypeName: "aws_s3_bucket_object",
			Name:     "Bucket Object",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
				ResourceType:        "BucketObject",
			},
		},
		{
			Factory:  dataSourceBucketObjects,
			TypeName: "aws_s3_bucket_objects",
			Name:     "Bucket Objects",
		},
		{
			Factory:  dataSourceBucketPolicy,
			TypeName: "aws_s3_bucket_policy",
			Name:     "Bucket Policy",
		},
		{
			Factory:  dataSourceObject,
			TypeName: "aws_s3_object",
			Name:     "Object",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
				ResourceType:        "Object",
			},
		},
		{
			Factory:  dataSourceObjects,
			TypeName: "aws_s3_objects",
			Name:     "Objects",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceBucket,
			TypeName: "aws_s3_bucket",
			Name:     "Bucket",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrBucket,
				ResourceType:        "Bucket",
			},
		},
		{
			Factory:  resourceBucketAccelerateConfiguration,
			TypeName: "aws_s3_bucket_accelerate_configuration",
			Name:     "Bucket Accelerate Configuration",
		},
		{
			Factory:  resourceBucketACL,
			TypeName: "aws_s3_bucket_acl",
			Name:     "Bucket ACL",
		},
		{
			Factory:  resourceBucketAnalyticsConfiguration,
			TypeName: "aws_s3_bucket_analytics_configuration",
			Name:     "Bucket Analytics Configuration",
		},
		{
			Factory:  resourceBucketCorsConfiguration,
			TypeName: "aws_s3_bucket_cors_configuration",
			Name:     "Bucket CORS Configuration",
		},
		{
			Factory:  resourceBucketIntelligentTieringConfiguration,
			TypeName: "aws_s3_bucket_intelligent_tiering_configuration",
			Name:     "Bucket Intelligent-Tiering Configuration",
		},
		{
			Factory:  resourceBucketInventory,
			TypeName: "aws_s3_bucket_inventory",
			Name:     "Bucket Inventory",
		},
		{
			Factory:  resourceBucketLifecycleConfiguration,
			TypeName: "aws_s3_bucket_lifecycle_configuration",
			Name:     "Bucket Lifecycle Configuration",
		},
		{
			Factory:  resourceBucketLogging,
			TypeName: "aws_s3_bucket_logging",
			Name:     "Bucket Logging",
		},
		{
			Factory:  resourceBucketMetric,
			TypeName: "aws_s3_bucket_metric",
			Name:     "Bucket Metric",
		},
		{
			Factory:  resourceBucketNotification,
			TypeName: "aws_s3_bucket_notification",
			Name:     "Bucket Notification",
		},
		{
			Factory:  resourceBucketObject,
			TypeName: "aws_s3_bucket_object",
			Name:     "Bucket Object",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
				ResourceType:        "BucketObject",
			},
		},
		{
			Factory:  resourceBucketObjectLockConfiguration,
			TypeName: "aws_s3_bucket_object_lock_configuration",
			Name:     "Bucket Object Lock Configuration",
		},
		{
			Factory:  resourceBucketOwnershipControls,
			TypeName: "aws_s3_bucket_ownership_controls",
			Name:     "Bucket Ownership Controls",
		},
		{
			Factory:  resourceBucketPolicy,
			TypeName: "aws_s3_bucket_policy",
			Name:     "Bucket Policy",
		},
		{
			Factory:  resourceBucketPublicAccessBlock,
			TypeName: "aws_s3_bucket_public_access_block",
			Name:     "Bucket Public Access Block",
		},
		{
			Factory:  resourceBucketReplicationConfiguration,
			TypeName: "aws_s3_bucket_replication_configuration",
			Name:     "Bucket Replication Configuration",
		},
		{
			Factory:  resourceBucketRequestPaymentConfiguration,
			TypeName: "aws_s3_bucket_request_payment_configuration",
			Name:     "Bucket Request Payment Configuration",
		},
		{
			Factory:  resourceBucketServerSideEncryptionConfiguration,
			TypeName: "aws_s3_bucket_server_side_encryption_configuration",
			Name:     "Bucket Server-side Encryption Configuration",
		},
		{
			Factory:  resourceBucketVersioning,
			TypeName: "aws_s3_bucket_versioning",
			Name:     "Bucket Versioning",
		},
		{
			Factory:  resourceBucketWebsiteConfiguration,
			TypeName: "aws_s3_bucket_website_configuration",
			Name:     "Bucket Website Configuration",
		},
		{
			Factory:  resourceObject,
			TypeName: "aws_s3_object",
			Name:     "Object",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
				ResourceType:        "Object",
			},
		},
		{
			Factory:  resourceObjectCopy,
			TypeName: "aws_s3_object_copy",
			Name:     "Object Copy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
				ResourceType:        "ObjectCopy",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.S3
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
