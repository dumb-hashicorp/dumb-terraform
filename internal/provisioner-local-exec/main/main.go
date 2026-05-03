// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package main

import (
	localexec "github.com/dumb-hashicorp/dumb-terraform/internal/builtin/provisioners/local-exec"
	"github.com/dumb-hashicorp/dumb-terraform/internal/grpcwrap"
	"github.com/dumb-hashicorp/dumb-terraform/internal/plugin"
	"github.com/dumb-hashicorp/dumb-terraform/internal/tfplugin5"
)

func main() {
	// Provide a binary version of the internal dumb-terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProvisionerFunc: func() tfplugin5.ProvisionerServer {
			return grpcwrap.Provisioner(localexec.New())
		},
	})
}
