// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"os"
	"reflect"
	"testing"
)

func TestPluginPath(t *testing.T) {
	td := testTempDir(t)
	defer os.RemoveAll(td)
	t.Chdir(td)

	pluginPath := []string{"a", "b", "c"}

	m := Meta{}
	if err := m.storePluginPath(pluginPath); err != nil {
		t.Fatal(err)
	}

	restoredPath, err := m.loadPluginPath()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(pluginPath, restoredPath) {
		t.Fatalf("expected plugin path %#v, got %#v", pluginPath, restoredPath)
	}
}

func TestInternalProviders(t *testing.T) {
	m := Meta{}
	internal := m.internalProviders()
	tfProvider, err := internal["dumb-terraform"]()
	if err != nil {
		t.Fatal(err)
	}

	schema := tfProvider.GetProviderSchema()
	_, found := schema.DataSources["dumb-terraform_remote_state"]
	if !found {
		t.Errorf("didn't find dumb-terraform_remote_state in internal \"dumb-terraform\" provider")
	}
}
