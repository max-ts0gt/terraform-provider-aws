// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package secretsmanager

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	awstypes "github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/max-ts0gt/terraform-provider-aws/internal/conns"
	"github.com/max-ts0gt/terraform-provider-aws/internal/logging"
	tftags "github.com/max-ts0gt/terraform-provider-aws/internal/tags"
	"github.com/max-ts0gt/terraform-provider-aws/internal/types/option"
	"github.com/max-ts0gt/terraform-provider-aws/names"
)

// []*SERVICE.Tag handling

// Tags returns secretsmanager service tags.
func Tags(tags tftags.KeyValueTags) []awstypes.Tag {
	result := make([]awstypes.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := awstypes.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from secretsmanager service tags.
func KeyValueTags(ctx context.Context, tags []awstypes.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.ToString(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// getTagsIn returns secretsmanager service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) []awstypes.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets secretsmanager service tags in Context.
func setTagsOut(ctx context.Context, tags []awstypes.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(KeyValueTags(ctx, tags))
	}
}

// updateTags updates secretsmanager service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn *secretsmanager.Client, identifier string, oldTagsMap, newTagsMap any, optFns ...func(*secretsmanager.Options)) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	ctx = tflog.SetField(ctx, logging.KeyResourceId, identifier)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.SecretsManager)
	if len(removedTags) > 0 {
		input := &secretsmanager.UntagResourceInput{
			SecretId: aws.String(identifier),
			TagKeys:  removedTags.Keys(),
		}

		_, err := conn.UntagResource(ctx, input, optFns...)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.SecretsManager)
	if len(updatedTags) > 0 {
		input := &secretsmanager.TagResourceInput{
			SecretId: aws.String(identifier),
			Tags:     Tags(updatedTags),
		}

		_, err := conn.TagResource(ctx, input, optFns...)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates secretsmanager service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return updateTags(ctx, meta.(*conns.AWSClient).SecretsManagerClient(ctx), identifier, oldTags, newTags)
}
