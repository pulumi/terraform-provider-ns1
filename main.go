package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/ns1-terraform/terraform-provider-ns1/ns1"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ns1.Provider})
}
