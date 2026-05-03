// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package plugin

import (
	"github.com/dumb-hashicorp/go-plugin"
	"github.com/dumb-hashicorp/dumb-terraform/internal/plugin6"
)

// VersionedPlugins includes both protocol 5 and 6 because this is the function
// called in providerFactory (command/meta_providers.go) to set up the initial
// plugin client config.
var VersionedPlugins = map[int]plugin.PluginSet{
	5: {
		"provider":    &GRPCProviderPlugin{},
		"provisioner": &GRPCProvisionerPlugin{},
	},
	6: {
		"provider": &plugin6.GRPCProviderPlugin{},
	},
}
