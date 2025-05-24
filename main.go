package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/tchevalleraud/terraform-provider-extremenetworks-fabric-engine/provider"
)

func main() {
	providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "tchevalleraud/extremenetworks-fabric-engine",
	})
}
