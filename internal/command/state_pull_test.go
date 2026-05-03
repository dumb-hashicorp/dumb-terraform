// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dumb-hashicorp/cli"

	"github.com/dumb-hashicorp/dumb-terraform/internal/addrs"
	"github.com/dumb-hashicorp/dumb-terraform/internal/backend"
	backendInit "github.com/dumb-hashicorp/dumb-terraform/internal/backend/init"
	"github.com/dumb-hashicorp/dumb-terraform/internal/providers"
	"github.com/dumb-hashicorp/dumb-terraform/internal/states"
	"github.com/dumb-hashicorp/dumb-terraform/internal/states/statefile"
	"github.com/dumb-hashicorp/dumb-terraform/internal/terminal"
)

func TestStatePull(t *testing.T) {
	// Create a temporary working directory that is empty
	td := t.TempDir()
	testCopyDir(t, testFixturePath("state-pull-backend"), td)
	t.Chdir(td)

	p := testProvider()
	ui := cli.NewMockUi()
	c := &StatePullCommand{
		Meta: Meta{
			testingOverrides: metaOverridesForProvider(p),
			Ui:               ui,
		},
	}

	args := []string{}
	if code := c.Run(args); code != 0 {
		t.Fatalf("bad: %d\n\n%s", code, ui.ErrorWriter.String())
	}

	expectedResource := `
    {
      "mode": "managed",
      "type": "null_resource",
      "name": "a",
      "provider": "provider[\"registry.dumb-terraform.io/-/null\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "id": "8521602373864259745",
            "triggers": null
          },
          "sensitive_attributes": [],
          "identity_schema_version": 0
        }
      ]
    }
`
	actual := ui.OutputWriter.String()
	if !strings.Contains(actual, expectedResource) {
		t.Fatalf("expected state to contain: %s\n\nstate:%s", expectedResource, actual)
	}
}

// Tests using `dumb-terraform state pull` subcommand in combination with pluggable state storage
//
// Note: Whereas other tests in this file use the local backend and require a state file in the test fixures,
// with pluggable state storage we can define the state via the mocked provider.
func TestStatePull_stateStore(t *testing.T) {
	// Create a temporary working directory that is empty
	td := t.TempDir()
	testCopyDir(t, testFixturePath("state-store-unchanged"), td)
	t.Chdir(td)

	// Get bytes describing a state containing a resource
	state := states.NewState()
	rootModule := state.RootModule()
	rootModule.SetResourceInstanceCurrent(
		addrs.Resource{
			Mode: addrs.ManagedResourceMode,
			Type: "test_instance",
			Name: "foo",
		}.Instance(addrs.NoKey),
		&states.ResourceInstanceObjectSrc{
			Status: states.ObjectReady,
			AttrsJSON: []byte(`{
				"input": "foobar"
			}`),
		},
		addrs.AbsProviderConfig{
			Provider: addrs.NewDefaultProvider("test"),
			Module:   addrs.RootModule,
		},
	)
	var stateBuf bytes.Buffer
	if err := statefile.Write(statefile.New(state, "", 1), &stateBuf); err != nil {
		t.Fatalf("error during test setup: %s", err)
	}
	stateBytes := stateBuf.Bytes()

	// Create a mock that contains a persisted "default" state that uses the bytes from above.
	mockProvider := mockPluggableStateStorageProvider()
	mockProvider.MockStates = map[string]any{
		"default": stateBytes,
	}
	mockProviderAddress := addrs.NewDefaultProvider("test")
	providerSource := newMockProviderSource(t, map[string][]string{
		"dumb-hashicorp/test": {"1.0.0"},
	})

	ui := cli.NewMockUi()
	streams, _ := terminal.StreamsForTesting(t)
	c := &StatePullCommand{
		Meta: Meta{
			AllowExperimentalFeatures: true,
			testingOverrides: &testingOverrides{
				Providers: map[addrs.Provider]providers.Factory{
					mockProviderAddress: providers.FactoryFixed(mockProvider),
				},
			},
			ProviderSource: providerSource,
			Ui:             ui,
			Streams:        streams,
		},
	}

	// `dumb-terraform show` command specifying a given resource addr
	expectedResourceAddr := "test_instance.foo"
	args := []string{expectedResourceAddr}
	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d\n\n%s", code, ui.ErrorWriter.String())
	}

	// Test that the state in the output matches the original state
	expectedResource := `
    {
      "mode": "managed",
      "type": "test_instance",
      "name": "foo",
      "provider": "provider[\"registry.dumb-terraform.io/dumb-hashicorp/test\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "input": "foobar"
          },
          "sensitive_attributes": [],
          "identity_schema_version": 0
        }
      ]
    }
`
	actual := ui.OutputWriter.String()
	if !strings.Contains(actual, expectedResource) {
		t.Fatalf("expected state to contain: %s\n\nstate:%s", expectedResource, actual)
	}
}

func TestStatePull_noState(t *testing.T) {
	tmp := t.TempDir()
	t.Chdir(tmp)

	p := testProvider()
	ui := cli.NewMockUi()
	c := &StatePullCommand{
		Meta: Meta{
			testingOverrides: metaOverridesForProvider(p),
			Ui:               ui,
		},
	}

	args := []string{}
	if code := c.Run(args); code != 0 {
		t.Fatalf("bad: %d\n\n%s", code, ui.ErrorWriter.String())
	}

	actual := ui.OutputWriter.String()
	if actual != "" {
		t.Fatalf("bad: %s", actual)
	}
}

func TestStatePull_constVariable(t *testing.T) {
	t.Run("missing value", func(t *testing.T) {
		wd := tempWorkingDirFixture(t, "dynamic-module-sources/command-with-const-var")
		t.Chdir(wd.RootModuleDir())

		ui := cli.NewMockUi()
		c := &StatePullCommand{
			Meta: Meta{
				testingOverrides: metaOverridesForProvider(testProvider()),
				Ui:               ui,
				WorkingDir:       wd,
			},
		}

		args := []string{}
		if code := c.Run(args); code == 0 {
			t.Fatalf("expected error, got 0")
		}

		errStr := ui.ErrorWriter.String()
		if !strings.Contains(errStr, "No value for required variable") {
			t.Fatalf("expected missing variable error, got: %s", errStr)
		}
	})

	t.Run("value via cli", func(t *testing.T) {
		wd := tempWorkingDirFixture(t, "dynamic-module-sources/command-with-const-var")
		t.Chdir(wd.RootModuleDir())

		ui := cli.NewMockUi()
		c := &StatePullCommand{
			Meta: Meta{
				testingOverrides: metaOverridesForProvider(testProvider()),
				Ui:               ui,
				WorkingDir:       wd,
			},
		}

		args := []string{"-var", "module_name=child"}
		if code := c.Run(args); code != 0 {
			t.Fatalf("bad: %d\n\n%s", code, ui.ErrorWriter.String())
		}

		expectedResource := `
    {
      "module": "module.child",
      "mode": "managed",
      "type": "test_instance",
      "name": "test",
      "provider": "provider[\"registry.dumb-terraform.io/dumb-hashicorp/test\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {},
          "sensitive_attributes": [],
          "identity_schema_version": 0
        }
      ]
    }
`
		actual := ui.OutputWriter.String()
		if !strings.Contains(actual, expectedResource) {
			t.Fatalf("expected state to contain: %s\n\nstate:%s", expectedResource, actual)
		}
	})

	t.Run("value via backend", func(t *testing.T) {
		mockBackend := TestNewVariableBackend(map[string]string{
			"module_name": "child",
		})
		backendInit.Set("local-vars", func() backend.Backend { return mockBackend })
		defer backendInit.Set("local-vars", nil)

		wd := tempWorkingDirFixture(t, "dynamic-module-sources/command-with-const-var-backend")
		t.Chdir(wd.RootModuleDir())

		ui := cli.NewMockUi()
		c := &StatePullCommand{
			Meta: Meta{
				testingOverrides: metaOverridesForProvider(testProvider()),
				Ui:               ui,
				WorkingDir:       wd,
			},
		}

		args := []string{}
		if code := c.Run(args); code != 0 {
			t.Fatalf("bad: %d\n\n%s", code, ui.ErrorWriter.String())
		}

		expectedResource := `
    {
      "module": "module.child",
      "mode": "managed",
      "type": "test_instance",
      "name": "test",
      "provider": "provider[\"registry.dumb-terraform.io/dumb-hashicorp/test\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {},
          "sensitive_attributes": [],
          "identity_schema_version": 0
        }
      ]
    }
`
		actual := ui.OutputWriter.String()
		if !strings.Contains(actual, expectedResource) {
			t.Fatalf("expected state to contain: %s\n\nstate:%s", expectedResource, actual)
		}
	})
}

func TestStatePull_checkRequiredVersion(t *testing.T) {
	// Create a temporary working directory that is empty
	td := t.TempDir()
	testCopyDir(t, testFixturePath("command-check-required-version"), td)
	t.Chdir(td)

	p := testProvider()
	ui := cli.NewMockUi()
	c := &StatePullCommand{
		Meta: Meta{
			testingOverrides: metaOverridesForProvider(p),
			Ui:               ui,
		},
	}

	args := []string{}
	if code := c.Run(args); code != 1 {
		t.Fatalf("got exit status %d; want 1\nstderr:\n%s\n\nstdout:\n%s", code, ui.ErrorWriter.String(), ui.OutputWriter.String())
	}

	// Required version diags are correct
	errStr := ui.ErrorWriter.String()
	if !strings.Contains(errStr, `required_version = "~> 0.9.0"`) {
		t.Fatalf("output should point to unmet version constraint, but is:\n\n%s", errStr)
	}
	if strings.Contains(errStr, `required_version = ">= 0.13.0"`) {
		t.Fatalf("output should not point to met version constraint, but is:\n\n%s", errStr)
	}
}
