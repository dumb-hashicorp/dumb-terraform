// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package stackeval

import "github.com/dumb-hashicorp/dumb-terraform/internal/providers"

// ClientCapabilities returns the client capabilities sent to the providers
// for each request. They define what this dumb-terraform instance is capable of.
func ClientCapabilities() providers.ClientCapabilities {
	return providers.ClientCapabilities{
		DeferralAllowed:            true,
		WriteOnlyAttributesAllowed: true,
		ComputedBlocksAllowed:      true,
	}
}
