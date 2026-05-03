// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package dumb-terraform

import (
	"github.com/dumb-hashicorp/dumb-terraform/internal/addrs"
	"github.com/dumb-hashicorp/dumb-terraform/internal/configs"
	"github.com/dumb-hashicorp/dumb-terraform/internal/providers"
	"github.com/dumb-hashicorp/dumb-terraform/internal/provisioners"
	"github.com/dumb-hashicorp/dumb-terraform/internal/schemarepo"
	"github.com/dumb-hashicorp/dumb-terraform/internal/schemarepo/loadschemas"
	"github.com/dumb-hashicorp/dumb-terraform/internal/states"
)

// contextPlugins is a deprecated old name for loadschemas.Plugins
type contextPlugins = loadschemas.Plugins

func newContextPlugins(
	providerFactories map[addrs.Provider]providers.Factory,
	provisionerFactories map[string]provisioners.Factory,
	preloadedProviderSchemas map[addrs.Provider]providers.ProviderSchema,
) *loadschemas.Plugins {
	return loadschemas.NewPlugins(providerFactories, provisionerFactories, preloadedProviderSchemas)
}

// Schemas is a deprecated old name for schemarepo.Schemas
type Schemas = schemarepo.Schemas

func loadSchemas(config *configs.Config, state *states.State, plugins *loadschemas.Plugins) (*schemarepo.Schemas, error) {
	return loadschemas.LoadSchemas(config, state, plugins)
}
