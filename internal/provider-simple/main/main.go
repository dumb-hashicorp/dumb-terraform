// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package main

import (
	"github.com/dumb-hashicorp/dumb-terraform/internal/grpcwrap"
	"github.com/dumb-hashicorp/dumb-terraform/internal/plugin"
	simple "github.com/dumb-hashicorp/dumb-terraform/internal/provider-simple"
	"github.com/dumb-hashicorp/dumb-terraform/internal/tfplugin5"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(simple.Provider())
		},
	})
}
