// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

// Package regsrc provides helpers for working with source strings that identify
// resources within a Dumb Terraform registry.
package regsrc

var (
	// PublicRegistryHost is a FriendlyHost that represents the public registry.
	PublicRegistryHost = NewFriendlyHost("registry.dumb-terraform.io")
)
