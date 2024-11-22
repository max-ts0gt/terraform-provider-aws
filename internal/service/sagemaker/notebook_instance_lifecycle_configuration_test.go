// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemaker_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/max-ts0gt/terraform-provider-aws/internal/acctest"
	"github.com/max-ts0gt/terraform-provider-aws/internal/conns"
	tfsagemaker "github.com/max-ts0gt/terraform-provider-aws/internal/service/sagemaker"
	"github.com/max-ts0gt/terraform-provider-aws/internal/tfresource"
	itypes "github.com/max-ts0gt/terraform-provider-aws/internal/types"
	"github.com/max-ts0gt/terraform-provider-aws/names"
)

func TestAccSageMakerNotebookInstanceLifecycleConfiguration_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var lifecycleConfig sagemaker.DescribeNotebookInstanceLifecycleConfigOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_notebook_instance_lifecycle_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SageMakerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckNotebookInstanceLifecycleConfigurationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceLifecycleConfigurationConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceLifecycleConfigurationExists(ctx, resourceName, &lifecycleConfig),

					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckNoResourceAttr(resourceName, "on_create"),
					resource.TestCheckNoResourceAttr(resourceName, "on_start"),
					acctest.CheckResourceAttrRegionalARN(resourceName, names.AttrARN, "sagemaker", fmt.Sprintf("notebook-instance-lifecycle-config/%s", rName)),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSageMakerNotebookInstanceLifecycleConfiguration_update(t *testing.T) {
	ctx := acctest.Context(t)
	var lifecycleConfig sagemaker.DescribeNotebookInstanceLifecycleConfigOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_notebook_instance_lifecycle_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SageMakerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckNotebookInstanceLifecycleConfigurationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceLifecycleConfigurationConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceLifecycleConfigurationExists(ctx, resourceName, &lifecycleConfig),

					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
				),
			},
			{
				Config: testAccNotebookInstanceLifecycleConfigurationConfig_update(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceLifecycleConfigurationExists(ctx, resourceName, &lifecycleConfig),

					resource.TestCheckResourceAttr(resourceName, "on_create", itypes.Base64EncodeOnce([]byte("echo bla"))),
					resource.TestCheckResourceAttr(resourceName, "on_start", itypes.Base64EncodeOnce([]byte("echo blub"))),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSageMakerNotebookInstanceLifecycleConfiguration_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var lifecycleConfig sagemaker.DescribeNotebookInstanceLifecycleConfigOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_notebook_instance_lifecycle_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SageMakerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckNotebookInstanceLifecycleConfigurationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebookInstanceLifecycleConfigurationConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotebookInstanceLifecycleConfigurationExists(ctx, resourceName, &lifecycleConfig),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceNotebookInstanceLifeCycleConfiguration(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckNotebookInstanceLifecycleConfigurationExists(ctx context.Context, resourceName string, lifecycleConfig *sagemaker.DescribeNotebookInstanceLifecycleConfigOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerClient(ctx)

		output, err := tfsagemaker.FindNotebookInstanceLifecycleConfigByName(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*lifecycleConfig = *output

		return nil
	}
}

func testAccCheckNotebookInstanceLifecycleConfigurationDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_sagemaker_notebook_instance_lifecycle_configuration" {
				continue
			}

			conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerClient(ctx)

			_, err := tfsagemaker.FindNotebookInstanceLifecycleConfigByName(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("SageMaker Notebook Instance Lifecycle Configuration %s still exists", rs.Primary.ID)
		}
		return nil
	}
}

func testAccNotebookInstanceLifecycleConfigurationConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance_lifecycle_configuration" "test" {
  name = %q
}
`, rName)
}

func testAccNotebookInstanceLifecycleConfigurationConfig_update(rName string) string {
	return fmt.Sprintf(`
resource "aws_sagemaker_notebook_instance_lifecycle_configuration" "test" {
  name      = %q
  on_create = base64encode("echo bla")
  on_start  = base64encode("echo blub")
}
`, rName)
}
