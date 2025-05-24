package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/tchevalleraud/terraform-provider-extremenetworks-fabric-engine/resources"
)

var _ provider.Provider = &extremeFabricProvider{}

type extremeFabricProvider struct{}

func New() provider.Provider {
	return &extremeFabricProvider{}
}

func (p *extremeFabricProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "extremenetworks-fabric-engine"
}

func (p *extremeFabricProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provider to manage Extreme Networks Fabric Engine devices via SSH.",
	}
}

func (p *extremeFabricProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// No shared provider config for now
}

func (p *extremeFabricProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewHostnameResource,
	}
}

func (p *extremeFabricProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}
