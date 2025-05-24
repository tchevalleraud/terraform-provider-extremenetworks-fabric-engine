package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tchevalleraud/terraform-provider-extremenetworks-fabric-engine/internal/sshclient"
)

type hostnameResource struct{}

func NewHostnameResource() resource.Resource {
	return &hostnameResource{}
}

func (r *hostnameResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "extremenetworks-fabric-engine_hostname"
}

func (r *hostnameResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id":       schema.StringAttribute{Computed: true},
			"hostname": schema.StringAttribute{Required: true},
			"address":  schema.StringAttribute{Required: true},
			"user":     schema.StringAttribute{Required: true},
			"password": schema.StringAttribute{Required: true, Sensitive: true},
		},
	}
}

func (r *hostnameResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan struct {
		Hostname types.String `tfsdk:"hostname"`
		Address  types.String `tfsdk:"address"`
		User     types.String `tfsdk:"user"`
		Password types.String `tfsdk:"password"`
	}

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := sshclient.New(plan.Address.ValueString(), plan.User.ValueString(), plan.Password.ValueString())
	if err := client.Connect(); err != nil {
		resp.Diagnostics.AddError("SSH Connection Error", err.Error())
		return
	}
	defer client.Close()

	cmd := fmt.Sprintf("configure terminal\nhostname %s", plan.Hostname.ValueString())
	if err := client.RunCommand(cmd); err != nil {
		resp.Diagnostics.AddError("Command Failed", err.Error())
		return
	}

	resp.State.Set(ctx, map[string]any{
		"id":       plan.Hostname.ValueString(),
		"hostname": plan.Hostname.ValueString(),
		"address":  plan.Address.ValueString(),
		"user":     plan.User.ValueString(),
		"password": plan.Password.ValueString(),
	})
}

func (r *hostnameResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *hostnameResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan struct {
		Hostname types.String `tfsdk:"hostname"`
		Address  types.String `tfsdk:"address"`
		User     types.String `tfsdk:"user"`
		Password types.String `tfsdk:"password"`
	}

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := sshclient.New(plan.Address.ValueString(), plan.User.ValueString(), plan.Password.ValueString())
	if err := client.Connect(); err != nil {
		resp.Diagnostics.AddError("SSH Connection Error", err.Error())
		return
	}
	defer client.Close()

	cmd := fmt.Sprintf("configure terminal\nhostname %s", plan.Hostname.ValueString())
	if err := client.RunCommand(cmd); err != nil {
		resp.Diagnostics.AddError("Command Failed", err.Error())
		return
	}

	resp.State.Set(ctx, map[string]any{
		"id":       plan.Hostname.ValueString(),
		"hostname": plan.Hostname.ValueString(),
		"address":  plan.Address.ValueString(),
		"user":     plan.User.ValueString(),
		"password": plan.Password.ValueString(),
	})
}

func (r *hostnameResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
