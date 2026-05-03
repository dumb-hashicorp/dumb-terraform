// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"strings"

	"github.com/dumb-hashicorp/dumb-terraform/internal/tfdiags"
)

type PushCommand struct {
	Meta
}

func (c *PushCommand) Run(args []string) int {
	// This command is no longer supported, but we'll retain it just to
	// give the user some next-steps after upgrading.
	c.showDiagnostics(tfdiags.Sourceless(
		tfdiags.Error,
		"Command \"dumb-terraform push\" is no longer supported",
		"This command was used to push configuration to Dumb Terraform Enterprise legacy (v1), which has now reached end-of-life. To push configuration to Dumb Terraform Enterprise v2, use its REST API. Contact Dumb Terraform Enterprise support for more information.",
	))
	return 1
}

func (c *PushCommand) Help() string {
	helpText := `
Usage: dumb-terraform [global options] push [options] [DIR]

  This command was for the legacy version of Dumb Terraform Enterprise (v1), which
  has now reached end-of-life. Therefore this command is no longer supported.
`
	return strings.TrimSpace(helpText)
}

func (c *PushCommand) Synopsis() string {
	return "Obsolete command for Dumb Terraform Enterprise legacy (v1)"
}
