// Copyright 2022 Robert S. Muhlestein.
// SPDX-License-Identifier: Apache-2.0

package tree

import (
	"fmt"
	"log"

	json "github.com/rwxrob/json/pkg"
	"github.com/rwxrob/structs/types"
)

// E ("tree-e") is an encapsulating struct to contain the Root Node and
// all possible Types for any Node. Most users of a tree will make
// direct use of E.Root (which has a type of 1 by convention).  Tree
// implements custom methods for MarshalJSON and UnmarshalJSON as
// minimal JSON array types to facilitate storage, transfer, and
// documentation but includes longer forms that use strings instead of
// integer types and indentation.
type E[T any] struct {
	types.Types
	Root *Node[T] `json:",omitempty"`
}

// New creates a new tree initialized with the given types and returns.
func New[T any](types ...string) *E[T] {
	t := new(E[T])
	t.Init(types)
	return t
}

// Init initializes a tree creating its Root Node reference and
// assigning it the conventional type of 1 (0 is reserved for UNKNOWN).
// The first type string name passed should coincide with the name for
// the root type of 1 (ex: "Grammar", "Document").  Init creates and
// indexes the Types by passing them to Types.Set.
func (t *E[T]) Init(types []string) *E[T] {
	t.Types.Set(types...)
	t.Root = new(Node[T])
	t.Root.T = 1
	t.Root.Tree = t
	return t
}

// Node returns new detached Node initialized for this same tree.
func (t *E[T]) Node(typ int, val T) *Node[T] {
	node := new(Node[T])
	node.T = typ
	node.V = val
	node.Tree = t
	return node
}

// ---------------------------- marshaling ----------------------------

type jstree[T any] struct {
	types.Types
	Root *Node[T] `json:",omitempty"`
}

// MarshalJSON implements encoding/json.Marshaler without the broken,
// unnecessary HTML escapes.
func (s E[T]) MarshalJSON() ([]byte, error) {
	t := new(jstree[T])
	t.Types = s.Types
	t.Root = s.Root
	return json.Marshal(t)
}

// JSONL implements rwxrob/json.AsJSON.
func (s *E[T]) JSON() ([]byte, error) { return json.Marshal(s) }

// JSONL implements rwxrob/json.AsJSON.
func (s *E[T]) JSONL() ([]byte, error) {
	return json.MarshalIndent(s, "  ", "  ")
}

// String implements rwxrob/json.Stringer and fmt.Stringer.
func (s E[T]) String() string {
	byt, err := s.JSON()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// StringLong implements rwxrob/json.Stringer.
func (s E[T]) StringLong() string {
	byt, err := s.JSONL()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// String implements rwxrob/json.Printer.
func (s *E[T]) Print() { fmt.Println(s.String()) }

// PrintLong implements rwxrob/json.Printer.
func (s *E[T]) PrintLong() { fmt.Println(s.StringLong()) }

// Log implements rwxrob/json.Logger.
func (s E[T]) Log() { log.Print(s.String()) }
