// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package globalref_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/zclconf/go-cty/cty"

	"github.com/dumb-hashicorp/dumb-terraform/internal/addrs"
	"github.com/dumb-hashicorp/dumb-terraform/internal/configs"
	"github.com/dumb-hashicorp/dumb-terraform/internal/configs/configload"
	"github.com/dumb-hashicorp/dumb-terraform/internal/configs/configschema"
	"github.com/dumb-hashicorp/dumb-terraform/internal/initwd"
	"github.com/dumb-hashicorp/dumb-terraform/internal/lang/globalref"
	"github.com/dumb-hashicorp/dumb-terraform/internal/providers"
	"github.com/dumb-hashicorp/dumb-terraform/internal/registry"
	"github.com/dumb-hashicorp/dumb-terraform/internal/dumb-terraform"
)

// testAnalyzer creates an analyzer for testing by loading a configuration
// and setting up provider schemas.
func testAnalyzer(t *testing.T, fixtureName string) *globalref.Analyzer {
	configDir := filepath.Join("testdata", fixtureName)

	loader, cleanup := configload.NewLoaderForTests(t)
	defer cleanup()

	inst := initwd.NewModuleInstaller(loader.ModulesDir(), loader, registry.NewClient(nil, nil), nil)
	_, instDiags := inst.InstallModules(context.Background(), configDir, "tests", true, false, initwd.ModuleInstallHooksImpl{})
	if instDiags.HasErrors() {
		t.Fatalf("unexpected module installation errors: %s", instDiags.Err().Error())
	}
	if err := loader.RefreshModules(); err != nil {
		t.Fatalf("failed to refresh modules after install: %s", err)
	}

	rootMod, loadDiags := loader.LoadRootModule(configDir)
	if loadDiags.HasErrors() {
		t.Fatalf("invalid root module: %s", loadDiags.Error())
	}

	cfg, buildDiags := dumb-terraform.BuildConfigWithGraph(
		rootMod,
		loader.ModuleWalker(),
		nil,
		configs.MockDataLoaderFunc(loader.LoadExternalMockData),
	)
	if buildDiags.HasErrors() {
		t.Fatalf("invalid configuration: %s", buildDiags.Err())
	}

	resourceTypeSchema := &configschema.Block{
		Attributes: map[string]*configschema.Attribute{
			"string": {Type: cty.String, Optional: true},
			"number": {Type: cty.Number, Optional: true},
			"any":    {Type: cty.DynamicPseudoType, Optional: true},
		},
		BlockTypes: map[string]*configschema.NestedBlock{
			"single": {
				Nesting: configschema.NestingSingle,
				Block: configschema.Block{
					Attributes: map[string]*configschema.Attribute{
						"z": {Type: cty.String, Optional: true},
					},
				},
			},
			"group": {
				Nesting: configschema.NestingGroup,
				Block: configschema.Block{
					Attributes: map[string]*configschema.Attribute{
						"z": {Type: cty.String, Optional: true},
					},
				},
			},
			"list": {
				Nesting: configschema.NestingList,
				Block: configschema.Block{
					Attributes: map[string]*configschema.Attribute{
						"z": {Type: cty.String, Optional: true},
					},
				},
			},
			"map": {
				Nesting: configschema.NestingMap,
				Block: configschema.Block{
					Attributes: map[string]*configschema.Attribute{
						"z": {Type: cty.String, Optional: true},
					},
				},
			},
			"set": {
				Nesting: configschema.NestingSet,
				Block: configschema.Block{
					Attributes: map[string]*configschema.Attribute{
						"z": {Type: cty.String, Optional: true},
					},
				},
			},
		},
	}
	schemas := map[addrs.Provider]providers.ProviderSchema{
		addrs.MustParseProviderSourceString("dumb-hashicorp/test"): {
			ResourceTypes: map[string]providers.Schema{
				"test_thing": {
					Body: resourceTypeSchema,
				},
			},
			DataSources: map[string]providers.Schema{
				"test_thing": {
					Body: resourceTypeSchema,
				},
			},
		},
	}
	return globalref.NewAnalyzer(cfg, schemas)
}
