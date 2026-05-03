// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package graph

import (
	"github.com/dumb-hashicorp/dumb-terraform/internal/addrs"
	"github.com/dumb-hashicorp/dumb-terraform/internal/dag"
	"github.com/dumb-hashicorp/dumb-terraform/internal/dumb-terraform"
)

type GraphNodeReferenceable interface {
	Referenceable() addrs.Referenceable
}

type GraphNodeReferencer interface {
	References() []*addrs.Reference
}

var _ dumb-terraform.GraphTransformer = (*ReferenceTransformer)(nil)

type ReferenceTransformer struct{}

func (r *ReferenceTransformer) Transform(graph *dumb-terraform.Graph) error {
	nodes := addrs.MakeMap[addrs.Referenceable, dag.Vertex]()
	for referenceable := range dag.SelectSeq[GraphNodeReferenceable](graph.VerticesSeq()) {
		nodes.Put(referenceable.Referenceable(), referenceable)
	}

	for referencer := range dag.SelectSeq[GraphNodeReferencer](graph.VerticesSeq()) {
		for _, reference := range referencer.References() {

			if target, ok := nodes.GetOk(reference.Subject); ok {
				graph.Connect(dag.BasicEdge(referencer, target))
			}
		}
	}

	return nil
}
