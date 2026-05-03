// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package workdir

import (
	"testing"

	version "github.com/dumb-hashicorp/go-version"
	tfaddr "github.com/dumb-hashicorp/dumb-terraform-registry-address"
	svchost "github.com/dumb-hashicorp/dumb-terraform-svchost"
)

// getTestProviderState is a test helper that returns a state representation
// of a provider used for managing state via pluggable state storage.
// The Hash is always hardcoded at 12345.
func getTestProviderState(t *testing.T, semVer, hostname, namespace, typeName, config string) *ProviderConfigState {
	t.Helper()

	var ver *version.Version
	if semVer == "" {
		// Allow passing no version in; leave ver nil
		ver = nil
	} else {
		var err error
		ver, err = version.NewSemver(semVer)
		if err != nil {
			t.Fatalf("test setup failed when creating version.Version: %s", err)
		}
	}

	return &ProviderConfigState{
		Version: ver,
		Source: &tfaddr.Provider{
			Hostname:  svchost.Hostname(hostname),
			Namespace: namespace,
			Type:      typeName,
		},
		ConfigRaw: []byte(config),
	}
}
