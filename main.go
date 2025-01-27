package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/vmware/terraform-provider-avi/avi"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: avi.Provider})
}
