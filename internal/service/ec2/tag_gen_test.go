// Code generated by internal/generate/tagresource/main.go; DO NOT EDIT.

package ec2_test

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/max-ts0gt/terraform-provider-aws/internal/acctest"
	"github.com/max-ts0gt/terraform-provider-aws/internal/conns"
	tfec2 "github.com/max-ts0gt/terraform-provider-aws/internal/service/ec2"
	tftags "github.com/max-ts0gt/terraform-provider-aws/internal/tags"
	"github.com/max-ts0gt/terraform-provider-aws/internal/tfresource"
	"github.com/max-ts0gt/terraform-provider-aws/names"
)

func testAccCheckTagDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Client(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_ec2_tag" {
				continue
			}

			identifier, key, err := tftags.GetResourceID(rs.Primary.ID)
			if err != nil {
				return err
			}

			_, err = tfec2.FindTag(ctx, conn, identifier, key)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("%s resource (%s) tag (%s) still exists", names.EC2, identifier, key)
		}

		return nil
	}
}

func testAccCheckTagExists(ctx context.Context, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		identifier, key, err := tftags.GetResourceID(rs.Primary.ID)
		if err != nil {
			return err
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Client(ctx)

		_, err = tfec2.FindTag(ctx, conn, identifier, key)

		return err
	}
}
