// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/fwdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	fwflex "github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	"github.com/hashicorp/terraform-provider-aws/internal/retry"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @FrameworkResource("aws_sesv2_account_dedicated_ip_warmup_attributes", name="Account Dedicated IP Warmup Attributes")
// @SingletonIdentity(identityDuplicateAttributes="id")
// @Testing(hasNoPreExistingResource=true)
// @Testing(generator=false)
// @Testing(serialize=true)
// @Testing(checkDestroyNoop=true)
func newAccountDedicatedIPWarmupAttributesResource(context.Context) (resource.ResourceWithConfigure, error) {
	r := &accountDedicatedIPWarmupAttributesResource{}

	return r, nil
}

type accountDedicatedIPWarmupAttributesResource struct {
	framework.ResourceWithModel[accountDedicatedIPWarmupAttributesResourceModel]
	framework.WithNoOpDelete
	framework.WithImportByIdentity
}

func (r *accountDedicatedIPWarmupAttributesResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"auto_warmup_enabled": schema.BoolAttribute{
				Required: true,
			},
			names.AttrID: framework.IDAttributeDeprecatedWithAlternate(path.Root(names.AttrRegion)),
		},
	}
}

func (r *accountDedicatedIPWarmupAttributesResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var data accountDedicatedIPWarmupAttributesResourceModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	conn := r.Meta().SESV2Client(ctx)

	if err := putAccountDedicatedIPWarmupAttributes(ctx, conn, data.AutoWarmupEnabled.ValueBool()); err != nil {
		response.Diagnostics.AddError("creating SESv2 Account Dedicated IP Warmup Attributes", err.Error())

		return
	}

	// Set values for unknowns.
	data.ID = fwflex.StringValueToFramework(ctx, r.Meta().Region(ctx))

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *accountDedicatedIPWarmupAttributesResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var data accountDedicatedIPWarmupAttributesResourceModel
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	conn := r.Meta().SESV2Client(ctx)

	autoWarmupEnabled, err := findAccountDedicatedIPAutoWarmupEnabled(ctx, conn)

	if retry.NotFound(err) {
		response.Diagnostics.Append(fwdiag.NewResourceNotFoundWarningDiagnostic(err))
		response.State.RemoveResource(ctx)

		return
	}

	if err != nil {
		response.Diagnostics.AddError("reading SESv2 Account Dedicated IP Warmup Attributes", err.Error())

		return
	}

	// GetAccount returns DedicatedIpAutoWarmupEnabled, which AutoFlex can't map to auto_warmup_enabled by name.
	data.AutoWarmupEnabled = types.BoolValue(autoWarmupEnabled)

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *accountDedicatedIPWarmupAttributesResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var data accountDedicatedIPWarmupAttributesResourceModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	conn := r.Meta().SESV2Client(ctx)

	if err := putAccountDedicatedIPWarmupAttributes(ctx, conn, data.AutoWarmupEnabled.ValueBool()); err != nil {
		response.Diagnostics.AddError("updating SESv2 Account Dedicated IP Warmup Attributes", err.Error())

		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func putAccountDedicatedIPWarmupAttributes(ctx context.Context, conn *sesv2.Client, autoWarmupEnabled bool) error {
	input := sesv2.PutAccountDedicatedIpWarmupAttributesInput{
		AutoWarmupEnabled: autoWarmupEnabled,
	}
	_, err := conn.PutAccountDedicatedIpWarmupAttributes(ctx, &input)

	return err
}

func findAccountDedicatedIPAutoWarmupEnabled(ctx context.Context, conn *sesv2.Client) (bool, error) {
	output, err := findAccount(ctx, conn)

	if err != nil {
		return false, err
	}

	return output.DedicatedIpAutoWarmupEnabled, nil
}

type accountDedicatedIPWarmupAttributesResourceModel struct {
	framework.WithRegionModel
	AutoWarmupEnabled types.Bool   `tfsdk:"auto_warmup_enabled"`
	ID                types.String `tfsdk:"id"`
}
