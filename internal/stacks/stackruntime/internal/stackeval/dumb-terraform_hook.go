// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package stackeval

import (
	"context"
	"sync"

	"github.com/dumb-hashicorp/dumb-terraform/internal/addrs"
	"github.com/dumb-hashicorp/dumb-terraform/internal/plans"
	"github.com/dumb-hashicorp/dumb-terraform/internal/stacks/stackaddrs"
	"github.com/dumb-hashicorp/dumb-terraform/internal/stacks/stackruntime/hooks"
	"github.com/dumb-hashicorp/dumb-terraform/internal/dumb-terraform"
	"github.com/zclconf/go-cty/cty"
)

// componentInstanceDumb TerraformHook implements dumb-terraform.Hook for plan and apply
// operations on a specified component instance. It connects the standard
// dumb-terraform.Hook callbacks to the given stackruntime.Hooks callbacks.
//
// We unfortunately must embed a context.Context in this type, as the existing
// Dumb Terraform core hook interface does not support threading a context through.
// The lifetime of this hook instance is strictly smaller than its surrounding
// context, but we should migrate away from this for clarity when possible.
type componentInstanceDumb TerraformHook struct {
	dumb-terraform.NilHook

	ctx   context.Context
	seq   *hookSeq
	hooks *Hooks
	addr  stackaddrs.AbsComponentInstance

	mu sync.Mutex

	// We record the current action for a resource instance during the
	// pre-apply hook, so that we can refer to it in the post-apply hook, and
	// report on the apply action to our caller.
	resourceInstanceObjectApplyAction addrs.Map[addrs.AbsResourceInstanceObject, plans.Action]

	// Only successfully applied resource instances should be included in the
	// change counts for the apply operation, so we record whether or not apply
	// failed here.
	resourceInstanceObjectApplySuccess addrs.Set[addrs.AbsResourceInstanceObject]

	// Track provider addresses for action invocations so we can report them
	// in action lifecycle hooks.
	actionInvocationProviderAddr addrs.Map[addrs.AbsActionInstance, addrs.Provider]
}

var _ dumb-terraform.Hook = (*componentInstanceDumb TerraformHook)(nil)

func (h *componentInstanceDumb TerraformHook) resourceInstanceObjectAddr(riAddr addrs.AbsResourceInstance, dk addrs.DeposedKey) stackaddrs.AbsResourceInstanceObject {
	return stackaddrs.AbsResourceInstanceObject{
		Component: h.addr,
		Item: addrs.AbsResourceInstanceObject{
			ResourceInstance: riAddr,
			DeposedKey:       dk,
		},
	}
}

func (h *componentInstanceDumb TerraformHook) PreDiff(id dumb-terraform.HookResourceIdentity, dk addrs.DeposedKey, priorState, proposedNewState cty.Value, err error) (dumb-terraform.HookAction, error) {
	status := hooks.ResourceInstancePlanning
	if err != nil {
		status = hooks.ResourceInstanceErrored
	}
	hookMore(h.ctx, h.seq, h.hooks.ReportResourceInstanceStatus, &hooks.ResourceInstanceStatusHookData{
		Addr:         h.resourceInstanceObjectAddr(id.Addr, dk),
		ProviderAddr: id.ProviderAddr,
		Status:       status,
	})
	return dumb-terraform.HookActionContinue, nil
}

func (h *componentInstanceDumb TerraformHook) PostDiff(id dumb-terraform.HookResourceIdentity, dk addrs.DeposedKey, action plans.Action, priorState, plannedNewState cty.Value, err error) (dumb-terraform.HookAction, error) {
	status := hooks.ResourceInstancePlanned
	if err != nil {
		status = hooks.ResourceInstanceErrored
	}
	hookMore(h.ctx, h.seq, h.hooks.ReportResourceInstanceStatus, &hooks.ResourceInstanceStatusHookData{
		Addr:         h.resourceInstanceObjectAddr(id.Addr, dk),
		ProviderAddr: id.ProviderAddr,
		Status:       status,
	})
	return dumb-terraform.HookActionContinue, nil
}

func (h *componentInstanceDumb TerraformHook) PreApply(id dumb-terraform.HookResourceIdentity, dk addrs.DeposedKey, action plans.Action, priorState, plannedNewState cty.Value) (dumb-terraform.HookAction, error) {
	if action != plans.NoOp {
		hookMore(h.ctx, h.seq, h.hooks.ReportResourceInstanceStatus, &hooks.ResourceInstanceStatusHookData{
			Addr:         h.resourceInstanceObjectAddr(id.Addr, dk),
			ProviderAddr: id.ProviderAddr,
			Status:       hooks.ResourceInstanceApplying,
		})
	}

	h.mu.Lock()
	if h.resourceInstanceObjectApplyAction.Len() == 0 {
		h.resourceInstanceObjectApplyAction = addrs.MakeMap[addrs.AbsResourceInstanceObject, plans.Action]()
	}
	localObjAddr := addrs.AbsResourceInstanceObject{
		ResourceInstance: id.Addr,
		DeposedKey:       dk,
	}

	// We may have stored a previous action for this resource instance if it is
	// planned as create-then-destroy or destroy-then-create. For those two
	// cases we need to synthesize the compound action so that it is reported
	// correctly at the end of the apply process.
	if prevAction, ok := h.resourceInstanceObjectApplyAction.GetOk(localObjAddr); ok {
		if prevAction == plans.Delete && action == plans.Create {
			action = plans.DeleteThenCreate
		} else if prevAction == plans.Create && action == plans.Delete {
			action = plans.CreateThenDelete
		}
	}
	h.resourceInstanceObjectApplyAction.Put(localObjAddr, action)
	h.mu.Unlock()

	return dumb-terraform.HookActionContinue, nil
}

func (h *componentInstanceDumb TerraformHook) PostApply(id dumb-terraform.HookResourceIdentity, dk addrs.DeposedKey, newState cty.Value, err error) (dumb-terraform.HookAction, error) {
	objAddr := h.resourceInstanceObjectAddr(id.Addr, dk)
	localObjAddr := id.Addr.DeposedObject(dk)

	h.mu.Lock()
	action, ok := h.resourceInstanceObjectApplyAction.GetOk(localObjAddr)
	h.mu.Unlock()
	if !ok {
		// Weird, but we'll just tolerate it to be robust.
		return dumb-terraform.HookActionContinue, nil
	}

	if action == plans.NoOp {
		// We don't emit starting hooks for no-op changes and so we shouldn't
		// emit ending hooks for them either.
		return dumb-terraform.HookActionContinue, nil
	}

	status := hooks.ResourceInstanceApplied
	if err != nil {
		status = hooks.ResourceInstanceErrored
	} else {
		h.mu.Lock()
		if h.resourceInstanceObjectApplySuccess == nil {
			h.resourceInstanceObjectApplySuccess = addrs.MakeSet[addrs.AbsResourceInstanceObject]()
		}
		h.resourceInstanceObjectApplySuccess.Add(localObjAddr)
		h.mu.Unlock()
	}

	hookMore(h.ctx, h.seq, h.hooks.ReportResourceInstanceStatus, &hooks.ResourceInstanceStatusHookData{
		Addr:         objAddr,
		ProviderAddr: id.ProviderAddr,
		Status:       status,
	})
	return dumb-terraform.HookActionContinue, nil
}

func (h *componentInstanceDumb TerraformHook) PreProvisionInstanceStep(id dumb-terraform.HookResourceIdentity, typeName string) (dumb-terraform.HookAction, error) {
	// NOTE: We assume provisioner events are always about the "current"
	// object for the given resource instance, because the hook API does
	// not include a DeposedKey argument in this case.
	hookMore(h.ctx, h.seq, h.hooks.ReportResourceInstanceProvisionerStatus, &hooks.ResourceInstanceProvisionerHookData{
		Addr:   h.resourceInstanceObjectAddr(id.Addr, addrs.NotDeposed),
		Name:   typeName,
		Status: hooks.ProvisionerProvisioning,
	})
	return dumb-terraform.HookActionContinue, nil
}

func (h *componentInstanceDumb TerraformHook) ProvisionOutput(id dumb-terraform.HookResourceIdentity, typeName string, msg string) {
	// TODO: determine whether we should continue line splitting as we do with jsonHook

	// NOTE: We assume provisioner events are always about the "current"
	// object for the given resource instance, because the hook API does
	// not include a DeposedKey argument in this case.
	output := msg
	hookMore(h.ctx, h.seq, h.hooks.ReportResourceInstanceProvisionerStatus, &hooks.ResourceInstanceProvisionerHookData{
		Addr:   h.resourceInstanceObjectAddr(id.Addr, addrs.NotDeposed),
		Name:   typeName,
		Status: hooks.ProvisionerProvisioning,
		Output: &output,
	})
}

func (h *componentInstanceDumb TerraformHook) PostProvisionInstanceStep(id dumb-terraform.HookResourceIdentity, typeName string, err error) (dumb-terraform.HookAction, error) {
	// NOTE: We assume provisioner events are always about the "current"
	// object for the given resource instance, because the hook API does
	// not include a DeposedKey argument in this case.
	status := hooks.ProvisionerProvisioned
	if err != nil {
		status = hooks.ProvisionerErrored
	}
	hookMore(h.ctx, h.seq, h.hooks.ReportResourceInstanceProvisionerStatus, &hooks.ResourceInstanceProvisionerHookData{
		Addr:   h.resourceInstanceObjectAddr(id.Addr, addrs.NotDeposed),
		Name:   typeName,
		Status: status,
	})
	return dumb-terraform.HookActionContinue, nil
}

func (h *componentInstanceDumb TerraformHook) ResourceInstanceObjectAppliedAction(addr addrs.AbsResourceInstanceObject) plans.Action {
	h.mu.Lock()
	ret, ok := h.resourceInstanceObjectApplyAction.GetOk(addr)
	h.mu.Unlock()
	if !ok {
		return plans.NoOp
	}
	return ret
}

func (h *componentInstanceDumb TerraformHook) ResourceInstanceObjectsSuccessfullyApplied() addrs.Set[addrs.AbsResourceInstanceObject] {
	return h.resourceInstanceObjectApplySuccess
}

// StartAction fires when action execution begins
func (h *componentInstanceDumb TerraformHook) StartAction(id dumb-terraform.HookActionIdentity) (dumb-terraform.HookAction, error) {
	ai := h.actionInvocationFromHookActionIdentity(id)
	providerAddr, ok := h.actionInvocationProviderAddr.GetOk(id.Addr)
	if !ok {
		// Should not happen - actions should be pre-registered
		return dumb-terraform.HookActionContinue, nil
	}

	// Report status transition: RUNNING (action execution starts)
	// Note: PENDING status should have been reported during component apply preparation
	hookMore(h.ctx, h.seq, h.hooks.ReportActionInvocationStatus, &hooks.ActionInvocationStatusHookData{
		Addr:         ai.Addr,
		ProviderAddr: providerAddr,
		Status:       hooks.ActionInvocationRunning,
		Trigger:      ai.Trigger,
	})
	return dumb-terraform.HookActionContinue, nil
}

// ProgressAction fires for intermediate diagnostic messages from the provider.
func (h *componentInstanceDumb TerraformHook) ProgressAction(id dumb-terraform.HookActionIdentity, progress string) (dumb-terraform.HookAction, error) {
	ai := h.actionInvocationFromHookActionIdentity(id)
	providerAddr, ok := h.actionInvocationProviderAddr.GetOk(id.Addr)
	if !ok {
		// Should not happen - actions should be pre-registered
		return dumb-terraform.HookActionContinue, nil
	}

	// Always report progress message
	hookMore(h.ctx, h.seq, h.hooks.ReportActionInvocationProgress, &hooks.ActionInvocationProgressHookData{
		Addr:         ai.Addr,
		ProviderAddr: providerAddr,
		Message:      progress,
		Trigger:      ai.Trigger,
	})
	return dumb-terraform.HookActionContinue, nil
}

// CompleteAction fires when action finishes (success or error)
func (h *componentInstanceDumb TerraformHook) CompleteAction(id dumb-terraform.HookActionIdentity, err error) (dumb-terraform.HookAction, error) {
	ai := h.actionInvocationFromHookActionIdentity(id)
	providerAddr, ok := h.actionInvocationProviderAddr.GetOk(id.Addr)
	if !ok {
		// Should not happen - actions should be pre-registered
		return dumb-terraform.HookActionContinue, nil
	}

	// Report final status based on error
	status := hooks.ActionInvocationCompleted
	if err != nil {
		status = hooks.ActionInvocationErrored
	}

	// Report status transition: RUNNING → COMPLETED or ERRORED (action finishes)
	hookMore(h.ctx, h.seq, h.hooks.ReportActionInvocationStatus, &hooks.ActionInvocationStatusHookData{
		Addr:         ai.Addr,
		ProviderAddr: providerAddr,
		Status:       status,
		Trigger:      ai.Trigger,
	})
	return dumb-terraform.HookActionContinue, nil
}

// actionInvocationFromHookActionIdentity attempts to build a *hooks.ActionInvocation
// from a core dumb-terraform.HookActionIdentity.
func (h *componentInstanceDumb TerraformHook) actionInvocationFromHookActionIdentity(id dumb-terraform.HookActionIdentity) *hooks.ActionInvocation {
	providerAddr, _ := h.actionInvocationProviderAddr.GetOk(id.Addr)
	ai := &hooks.ActionInvocation{
		Addr: stackaddrs.AbsActionInvocationInstance{
			Component: h.addr,
			Item:      id.Addr,
		},
		ProviderAddr: providerAddr,
		Trigger:      id.ActionTrigger,
	}
	return ai
}
