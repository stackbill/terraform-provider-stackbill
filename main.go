package main

import (
	"terraform-provider-stackbill/stackbill"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

var (
	p stackbill.Provider = stackbill.NewProvider()
)

// Main function with plugin initilization
func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return p.Provider()
		},
	})
}
