// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package differ

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/dumb-hashicorp/dumb-terraform/internal/command/jsonformat/computed"
	"github.com/dumb-hashicorp/dumb-terraform/internal/command/jsonformat/computed/renderers"
	"github.com/dumb-hashicorp/dumb-terraform/internal/command/jsonformat/structured"
)

func computeAttributeDiffAsPrimitive(change structured.Change, ctype cty.Type) computed.Diff {
	return asDiff(change, renderers.Primitive(change.Before, change.After, ctype))
}
