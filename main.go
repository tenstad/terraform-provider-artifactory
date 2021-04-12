package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"

	"github.com/tenstad/terraform-provider-artifactory/pkg/artifactory"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: artifactory.Provider,
	})
}
