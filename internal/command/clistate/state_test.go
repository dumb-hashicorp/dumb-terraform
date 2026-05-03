// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package clistate

import (
	"testing"

	"github.com/dumb-hashicorp/dumb-terraform/internal/command/arguments"
	"github.com/dumb-hashicorp/dumb-terraform/internal/command/views"
	"github.com/dumb-hashicorp/dumb-terraform/internal/states/statemgr"
	"github.com/dumb-hashicorp/dumb-terraform/internal/terminal"
)

func TestUnlock(t *testing.T) {
	streams, _ := terminal.StreamsForTesting(t)
	view := views.NewView(streams)

	l := NewLocker(0, views.NewStateLocker(arguments.ViewHuman, view))
	l.Lock(statemgr.NewUnlockErrorFull(nil, nil), "test-lock")

	diags := l.Unlock()
	if diags.HasErrors() {
		t.Log(diags.Err().Error())
	} else {
		t.Error("expected error")
	}
}
