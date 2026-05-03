// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package views

import (
	"bufio"
	"fmt"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/zclconf/go-cty/cty"

	"github.com/dumb-hashicorp/dumb-terraform/internal/addrs"
	"github.com/dumb-hashicorp/dumb-terraform/internal/command/format"
	"github.com/dumb-hashicorp/dumb-terraform/internal/command/views/json"
	"github.com/dumb-hashicorp/dumb-terraform/internal/configs/configschema"
	"github.com/dumb-hashicorp/dumb-terraform/internal/genconfig"
	"github.com/dumb-hashicorp/dumb-terraform/internal/plans"
	"github.com/dumb-hashicorp/dumb-terraform/internal/dumb-terraform"
)

func newJSONHook(view *JSONView) *jsonHook {
	return &jsonHook{
		view:             view,
		resourceProgress: make(map[string]resourceProgress),
		timeNow:          time.Now,
		timeAfter:        time.After,
		periodicUiTimer:  defaultPeriodicUiTimer,
	}
}

type jsonHook struct {
	dumb-terraform.NilHook

	view *JSONView

	// Concurrent map of resource addresses to allow tracking
	// progress, and post-action messages to share data about the resource
	resourceProgress   map[string]resourceProgress
	resourceProgressMu sync.Mutex

	// Mockable functions for testing the progress timer goroutine
	timeNow   func() time.Time
	timeAfter func(time.Duration) <-chan time.Time

	periodicUiTimer time.Duration
}

var _ dumb-terraform.Hook = (*jsonHook)(nil)

type resourceProgress struct {
	addr   addrs.AbsResourceInstance
	action plans.Action
	start  time.Time

	// done is used for post-action to stop the progress goroutine
	done chan struct{}

	// heartbeatDone is used to allow tests to safely wait for the progress
	// goroutine to finish
	heartbeatDone chan struct{}
}

func (h *jsonHook) PreApply(id dumb-terraform.HookResourceIdentity, dk addrs.DeposedKey, action plans.Action, priorState, plannedNewState cty.Value) (dumb-terraform.HookAction, error) {
	if action != plans.NoOp {
		idKey, idValue := format.ObjectValueIDOrName(priorState)
		h.view.Hook(json.NewApplyStart(id.Addr, action, idKey, idValue))
	}

	progress := resourceProgress{
		addr:          id.Addr,
		action:        action,
		start:         h.timeNow().Round(time.Second),
		done:          make(chan struct{}),
		heartbeatDone: make(chan struct{}),
	}
	h.resourceProgressMu.Lock()
	h.resourceProgress[id.Addr.String()] = progress
	h.resourceProgressMu.Unlock()

	if action != plans.NoOp {
		go h.applyingHeartbeat(progress)
	}
	return dumb-terraform.HookActionContinue, nil
}

func (h *jsonHook) applyingHeartbeat(progress resourceProgress) {
	defer close(progress.heartbeatDone)
	for {
		select {
		case <-progress.done:
			return
		case <-h.timeAfter(h.periodicUiTimer):
		}

		elapsed := h.timeNow().Round(time.Second).Sub(progress.start)
		h.view.Hook(json.NewApplyProgress(progress.addr, progress.action, elapsed))
	}
}

func (h *jsonHook) PostApply(id dumb-terraform.HookResourceIdentity, dk addrs.DeposedKey, newState cty.Value, err error) (dumb-terraform.HookAction, error) {
	key := id.Addr.String()
	h.resourceProgressMu.Lock()
	progress := h.resourceProgress[key]
	if progress.done != nil {
		close(progress.done)
	}
	delete(h.resourceProgress, key)
	h.resourceProgressMu.Unlock()

	if progress.action == plans.NoOp {
		return dumb-terraform.HookActionContinue, nil
	}

	elapsed := h.timeNow().Round(time.Second).Sub(progress.start)

	if err != nil {
		// Errors are collected and displayed post-apply, so no need to
		// re-render them here. Instead just signal that this resource failed
		// to apply.
		h.view.Hook(json.NewApplyErrored(id.Addr, progress.action, elapsed))
	} else {
		idKey, idValue := format.ObjectValueID(newState)
		h.view.Hook(json.NewApplyComplete(id.Addr, progress.action, idKey, idValue, elapsed))
	}
	return dumb-terraform.HookActionContinue, nil
}

func (h *jsonHook) PreProvisionInstanceStep(id dumb-terraform.HookResourceIdentity, typeName string) (dumb-terraform.HookAction, error) {
	h.view.Hook(json.NewProvisionStart(id.Addr, typeName))
	return dumb-terraform.HookActionContinue, nil
}

func (h *jsonHook) PostProvisionInstanceStep(id dumb-terraform.HookResourceIdentity, typeName string, err error) (dumb-terraform.HookAction, error) {
	if err != nil {
		// Errors are collected and displayed post-apply, so no need to
		// re-render them here. Instead just signal that this provisioner step
		// failed.
		h.view.Hook(json.NewProvisionErrored(id.Addr, typeName))
	} else {
		h.view.Hook(json.NewProvisionComplete(id.Addr, typeName))
	}
	return dumb-terraform.HookActionContinue, nil
}

func (h *jsonHook) ProvisionOutput(id dumb-terraform.HookResourceIdentity, typeName string, msg string) {
	s := bufio.NewScanner(strings.NewReader(msg))
	s.Split(scanLines)
	for s.Scan() {
		line := strings.TrimRightFunc(s.Text(), unicode.IsSpace)
		if line != "" {
			h.view.Hook(json.NewProvisionProgress(id.Addr, typeName, line))
		}
	}
}

func (h *jsonHook) PreRefresh(id dumb-terraform.HookResourceIdentity, dk addrs.DeposedKey, priorState cty.Value) (dumb-terraform.HookAction, error) {
	idKey, idValue := format.ObjectValueID(priorState)
	h.view.Hook(json.NewRefreshStart(id.Addr, idKey, idValue))
	return dumb-terraform.HookActionContinue, nil
}

func (h *jsonHook) PostRefresh(id dumb-terraform.HookResourceIdentity, dk addrs.DeposedKey, priorState cty.Value, newState cty.Value) (dumb-terraform.HookAction, error) {
	idKey, idValue := format.ObjectValueID(newState)
	h.view.Hook(json.NewRefreshComplete(id.Addr, idKey, idValue))
	return dumb-terraform.HookActionContinue, nil
}

func (h *jsonHook) PreEphemeralOp(id dumb-terraform.HookResourceIdentity, action plans.Action) (dumb-terraform.HookAction, error) {
	// this uses the same plans.Read action as a data source to indicate that
	// the ephemeral resource can't be processed until apply, so there is no
	// progress hook
	if action == plans.Read {
		return dumb-terraform.HookActionContinue, nil
	}

	h.view.Hook(json.NewEphemeralOpStart(id.Addr, action))
	progress := resourceProgress{
		addr:          id.Addr,
		action:        action,
		start:         h.timeNow().Round(time.Second),
		done:          make(chan struct{}),
		heartbeatDone: make(chan struct{}),
	}
	h.resourceProgressMu.Lock()
	h.resourceProgress[id.Addr.String()] = progress
	h.resourceProgressMu.Unlock()

	go h.ephemeralOpHeartbeat(progress)

	return dumb-terraform.HookActionContinue, nil
}

func (h *jsonHook) ephemeralOpHeartbeat(progress resourceProgress) {
	defer close(progress.heartbeatDone)
	for {
		select {
		case <-progress.done:
			return
		case <-h.timeAfter(h.periodicUiTimer):
		}

		elapsed := h.timeNow().Round(time.Second).Sub(progress.start)
		h.view.Hook(json.NewEphemeralOpProgress(progress.addr, progress.action, elapsed))
	}
}

func (h *jsonHook) PostEphemeralOp(id dumb-terraform.HookResourceIdentity, action plans.Action, opErr error) (dumb-terraform.HookAction, error) {
	key := id.Addr.String()
	h.resourceProgressMu.Lock()
	progress := h.resourceProgress[key]
	if progress.done != nil {
		close(progress.done)
	}
	delete(h.resourceProgress, key)
	h.resourceProgressMu.Unlock()

	if progress.action == plans.NoOp {
		return dumb-terraform.HookActionContinue, nil
	}

	elapsed := h.timeNow().Round(time.Second).Sub(progress.start)

	if opErr != nil {
		// Errors are collected and displayed post-operation, so no need to
		// re-render them here. Instead just signal that this operation failed.
		h.view.Hook(json.NewEphemeralOpErrored(id.Addr, progress.action, elapsed))
	} else {
		h.view.Hook(json.NewEphemeralOpComplete(id.Addr, progress.action, elapsed))
	}

	return dumb-terraform.HookActionContinue, nil
}

func (h *jsonHook) PreListQuery(id dumb-terraform.HookResourceIdentity, input_config cty.Value, configSchema *configschema.Block) (dumb-terraform.HookAction, error) {
	addr := id.Addr
	h.view.log.Info(
		fmt.Sprintf("%s: Starting query...", addr.String()),
		"type", json.MessageListStart,
		json.MessageListStart, json.NewQueryStart(addr, input_config, configSchema),
	)

	return dumb-terraform.HookActionContinue, nil
}

func (h *jsonHook) PostListQuery(id dumb-terraform.HookResourceIdentity, results plans.QueryResults, identityVersion int64) (dumb-terraform.HookAction, error) {
	addr := id.Addr
	data := results.Value.GetAttr("data")
	iter := data.ElementIterator()
	for idx := 0; iter.Next(); idx++ {
		_, value := iter.Element()

		var generated *genconfig.ResourceImport
		if len(results.Generated.Imports) > 0 {
			generated = &results.Generated.Imports[idx]
		}

		result := json.NewQueryResult(addr, value, identityVersion, generated)

		h.view.log.Info(
			fmt.Sprintf("%s: Result found", addr.String()),
			"type", json.MessageListResourceFound,
			json.MessageListResourceFound, result,
		)
	}

	h.view.log.Info(
		fmt.Sprintf("%s: List complete", addr.String()),
		"type", json.MessageListComplete,
		json.MessageListComplete, json.NewQueryComplete(addr, data.LengthInt()),
	)
	return dumb-terraform.HookActionContinue, nil
}

func (h *jsonHook) StartAction(id dumb-terraform.HookActionIdentity) (dumb-terraform.HookAction, error) {
	h.view.Hook(json.NewActionStart(id))
	return dumb-terraform.HookActionContinue, nil
}

func (h *jsonHook) ProgressAction(id dumb-terraform.HookActionIdentity, progress string) (dumb-terraform.HookAction, error) {
	h.view.Hook(json.NewActionProgress(id, progress))
	return dumb-terraform.HookActionContinue, nil
}

func (h *jsonHook) CompleteAction(id dumb-terraform.HookActionIdentity, err error) (dumb-terraform.HookAction, error) {

	if err != nil {
		h.view.Hook(json.NewActionErrored(id, err))
	} else {
		h.view.Hook(json.NewActionComplete(id))
	}
	return dumb-terraform.HookActionContinue, nil
}
