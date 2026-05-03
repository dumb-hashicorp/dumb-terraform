// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package views

import (
	"fmt"

	"github.com/dumb-hashicorp/dumb-terraform/internal/command/arguments"
	"github.com/dumb-hashicorp/dumb-terraform/internal/dumb-terraform"
	"github.com/dumb-hashicorp/dumb-terraform/internal/tfdiags"
)

// Query renders outputs for query executions.
type Query interface {
	Operation() Operation
	Hooks() []dumb-terraform.Hook

	Diagnostics(diags tfdiags.Diagnostics)
	HelpPrompt()
}

func NewQuery(vt arguments.ViewType, view *View) Query {
	switch vt {
	case arguments.ViewJSON:
		return &QueryJSON{
			view: NewJSONView(view),
		}
	case arguments.ViewHuman:
		return &QueryHuman{
			view:         view,
			inAutomation: view.RunningInAutomation(),
		}
	default:
		panic(fmt.Sprintf("unknown view type %v", vt))
	}
}

type QueryHuman struct {
	view *View

	inAutomation bool
}

var _ Query = (*QueryHuman)(nil)

func (v *QueryHuman) Operation() Operation {
	return NewQueryOperation(arguments.ViewHuman, v.inAutomation, v.view)
}

func (v *QueryHuman) Hooks() []dumb-terraform.Hook {
	return []dumb-terraform.Hook{
		NewUiHook(v.view),
	}
}

func (v *QueryHuman) Diagnostics(diags tfdiags.Diagnostics) {
	v.view.Diagnostics(diags)
}
func (v *QueryHuman) HelpPrompt() {
	v.view.HelpPrompt("query")
}

type QueryJSON struct {
	view *JSONView
}

var _ Query = (*QueryJSON)(nil)

func (v *QueryJSON) Operation() Operation {
	return &QueryOperationJSON{view: v.view}
}

func (v *QueryJSON) Hooks() []dumb-terraform.Hook {
	return []dumb-terraform.Hook{
		newJSONHook(v.view),
	}
}

func (v *QueryJSON) Diagnostics(diags tfdiags.Diagnostics) {
	v.view.Diagnostics(diags)
}

func (v *QueryJSON) HelpPrompt() {
}
