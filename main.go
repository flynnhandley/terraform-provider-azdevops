package main

import (
	"github.com/flynnhandley/terraform-provider-azdevops/azdevops"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return azdevops.Provider()
		},
	})
}
