// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package sesv2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	tfsesv2 "github.com/hashicorp/terraform-provider-aws/internal/service/sesv2"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccSESV2AccountDedicatedIPWarmupAttributes_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]func(t *testing.T){
		acctest.CtBasic: testAccAccountDedicatedIPWarmupAttributes_basic,
		"update":        testAccAccountDedicatedIPWarmupAttributes_update,
		"Identity":      testAccSESV2AccountDedicatedIPWarmupAttributes_identitySerial,
	}

	acctest.RunSerialTests1Level(t, testCases, 0)
}

func testAccAccountDedicatedIPWarmupAttributes_basic(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_sesv2_account_dedicated_ip_warmup_attributes.test"

	acctest.Test(ctx, t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SESV2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             acctest.CheckDestroyNoop,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/AccountDedicatedIPWarmupAttributes/basic/"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountDedicatedIPWarmupAttributesExists(ctx, t, resourceName),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_warmup_enabled"), knownvalue.Bool(false)),
				},
			},
			{
				ConfigDirectory:                      config.StaticDirectory("testdata/AccountDedicatedIPWarmupAttributes/basic/"),
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: names.AttrRegion,
				ImportStateIdFunc:                    acctest.AttrImportStateIdFunc(resourceName, names.AttrRegion),
			},
		},
	})
}

func testAccAccountDedicatedIPWarmupAttributes_update(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_sesv2_account_dedicated_ip_warmup_attributes.test"

	acctest.Test(ctx, t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SESV2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             acctest.CheckDestroyNoop,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/AccountDedicatedIPWarmupAttributes/basic/"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountDedicatedIPWarmupAttributesExists(ctx, t, resourceName),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_warmup_enabled"), knownvalue.Bool(false)),
				},
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/AccountDedicatedIPWarmupAttributes/enabled/"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountDedicatedIPWarmupAttributesExists(ctx, t, resourceName),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_warmup_enabled"), knownvalue.Bool(true)),
				},
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/AccountDedicatedIPWarmupAttributes/basic/"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountDedicatedIPWarmupAttributesExists(ctx, t, resourceName),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("auto_warmup_enabled"), knownvalue.Bool(false)),
				},
			},
		},
	})
}

func testAccCheckAccountDedicatedIPWarmupAttributesExists(ctx context.Context, t *testing.T, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.ProviderMeta(ctx, t).SESV2Client(ctx)

		_, err := tfsesv2.FindAccountDedicatedIPAutoWarmupEnabled(ctx, conn)

		return err
	}
}
