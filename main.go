package main

import (
	"context"
	"flag"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"log"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debuggable", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/ochk/ochk",
			&plugin.ServeOpts{
				ProviderFunc: ochk.Provider,
			})
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		plugin.Serve(
			&plugin.ServeOpts{
				ProviderFunc: ochk.Provider,
			})
	}
}
