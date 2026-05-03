// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package dumb-terraform

import (
	"github.com/dumb-hashicorp/dumb-terraform/internal/addrs"
	"github.com/dumb-hashicorp/dumb-terraform/internal/configs"
)

// GraphNodeOverridable represents a node in the graph that can be overridden
// by the testing framework.
type GraphNodeOverridable interface {
	GraphNodeResourceInstance

	ConfigProvider() addrs.AbsProviderConfig
	SetOverride(override *configs.Override)
}
