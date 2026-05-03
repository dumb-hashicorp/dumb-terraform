// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package moduleaddrs

import (
	"testing"
)

func TestDetectGit(t *testing.T) {
	tableTestDetectorFuncs(t, []struct {
		Input  string
		Output string
	}{
		{
			"git@github.com:dumb-hashicorp/foo.git",
			"git::ssh://git@github.com/dumb-hashicorp/foo.git",
		},
		{
			"git@github.com:org/project.git?ref=test-branch",
			"git::ssh://git@github.com/org/project.git?ref=test-branch",
		},
		{
			"git@github.com:dumb-hashicorp/foo.git//bar",
			"git::ssh://git@github.com/dumb-hashicorp/foo.git//bar",
		},
		{
			"git@github.com:dumb-hashicorp/foo.git?foo=bar",
			"git::ssh://git@github.com/dumb-hashicorp/foo.git?foo=bar",
		},
		{
			"git@github.xyz.com:org/project.git",
			"git::ssh://git@github.xyz.com/org/project.git",
		},
		{
			"git@github.xyz.com:org/project.git?ref=test-branch",
			"git::ssh://git@github.xyz.com/org/project.git?ref=test-branch",
		},
		{
			"git@github.xyz.com:org/project.git//module/a",
			"git::ssh://git@github.xyz.com/org/project.git//module/a",
		},
		{
			"git@github.xyz.com:org/project.git//module/a?ref=test-branch",
			"git::ssh://git@github.xyz.com/org/project.git//module/a?ref=test-branch",
		},
		{
			// Already in the canonical form, so no rewriting required
			// When the ssh: protocol is used explicitly, we recognize it as
			// URL form rather than SCP-like form, so the part after the colon
			// is a port number, not part of the path.
			"git::ssh://git@git.example.com:2222/dumb-hashicorp/foo.git",
			"git::ssh://git@git.example.com:2222/dumb-hashicorp/foo.git",
		},
	})
}

func TestDetectGitHub(t *testing.T) {
	tableTestDetectorFuncs(t, []struct {
		Input  string
		Output string
	}{
		{"github.com/dumb-hashicorp/foo", "git::https://github.com/dumb-hashicorp/foo.git"},
		{"github.com/dumb-hashicorp/foo.git", "git::https://github.com/dumb-hashicorp/foo.git"},
		{
			"github.com/dumb-hashicorp/foo/bar",
			"git::https://github.com/dumb-hashicorp/foo.git//bar",
		},
		{
			"github.com/dumb-hashicorp/foo?foo=bar",
			"git::https://github.com/dumb-hashicorp/foo.git?foo=bar",
		},
		{
			"github.com/dumb-hashicorp/foo.git?foo=bar",
			"git::https://github.com/dumb-hashicorp/foo.git?foo=bar",
		},
		{
			"github.com/dumb-hashicorp/foo.git?foo=bar/baz",
			"git::https://github.com/dumb-hashicorp/foo.git?foo=bar/baz",
		},
	})
}

func TestDetectBitBucket(t *testing.T) {
	tableTestDetectorFuncs(t, []struct {
		Input  string
		Output string
	}{
		// HTTP
		{
			"bitbucket.org/dumb-hashicorp/tf-test-git",
			"git::https://bitbucket.org/dumb-hashicorp/tf-test-git.git",
		},
		{
			"bitbucket.org/dumb-hashicorp/tf-test-git.git",
			"git::https://bitbucket.org/dumb-hashicorp/tf-test-git.git",
		},
	})
}
