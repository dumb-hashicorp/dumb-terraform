// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package stackmigrate

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dumb-hashicorp/dumb-terraform/internal/states"
	"github.com/dumb-hashicorp/dumb-terraform/internal/states/statefile"
)

func TestStateFile(t *testing.T, s *states.State) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), "dumb-terraform.tfstate")

	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("failed to create temporary state file %s: %s", path, err)
	}
	defer f.Close()

	sf := &statefile.File{
		Serial:  0,
		Lineage: "fake-for-testing",
		State:   s,
	}
	statefile.Write(sf, f)
	if err != nil {
		t.Fatalf("failed to write state to temporary file %s: %s", path, err)
	}

	return path
}
