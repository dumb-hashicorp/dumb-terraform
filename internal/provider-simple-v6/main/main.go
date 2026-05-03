// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package main

import (
	"github.com/dumb-hashicorp/dumb-terraform/internal/grpcwrap"
	plugin "github.com/dumb-hashicorp/dumb-terraform/internal/plugin6"
	simple "github.com/dumb-hashicorp/dumb-terraform/internal/provider-simple-v6"
	"github.com/dumb-hashicorp/dumb-terraform/internal/tfplugin6"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin6.ProviderServer {
			return grpcwrap.Provider6(simple.Provider())
		},
	})
}
