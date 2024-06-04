// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grafana

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/grafana"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func statusWorkspaceStatus(ctx context.Context, conn *grafana.Client, id string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := FindWorkspaceByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, string(output.Status), nil
	}
}

func statusWorkspaceSAMLConfiguration(ctx context.Context, conn *grafana.Client, id string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := FindSamlConfigurationByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, string(output.Status), nil
	}
}
