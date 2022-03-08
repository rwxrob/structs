// Copyright 2022 Robert S. Muhlestein.
// SPDX-License-Identifier: Apache-2.0

package structs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rwxrob/fn/each"
	"github.com/rwxrob/to"
)

// Tree is an encapsulating struct to contain the Root Node and all
// possible Types for any Node.Most users of a Tree will make direct use
// of Tree.Root (which can safely be swapped out for a root Node
// reference at any time). Tree implements the rwxrob/json/AsJSON interface
// and uses custom methods for MarshalJSON and UnmarshalJSON to
// facilitate storage, transfer, and documentation.
type Tree[T any] struct {
	Root     *Node[T]       `json:",omitempty"`
	Types    []string       `json:",omitempty"`
	TypesMap map[string]int `json:",omitempty"`
}

// New initializes a tree creating its Root Node reference and assigning
// it the type of 1 (0 is reserved for UNKNOWN). Init then assigns its
// Types and indexes the TypesMap.
func (t *Tree[T]) Init(types []string) *Tree[T] {
	t.Types = []string{"UNKNOWN"}
	t.Types = append(t.Types, types...)
	t.TypesMap = map[string]int{"UNKNOWN": 0}
	for n, v := range types {
		t.TypesMap[v] = n + 1
	}
	t.Root = new(Node[T])
	t.Root.T = 1
	t.Root.Tree = t
	return t
}

// Node returns new detached Node initialized for this same Tree. Either
// an integer or string type may be passed and if none are passed the
// UNKNOWN (0) type will be assigned.
func (t *Tree[T]) Node(typ ...any) *Node[T] {
	node := new(Node[T])
	switch len(typ) {
	case 2:
		node.V = typ[1].(T)
		fallthrough
	case 1:
		node.T = typ[0].(int)
	}
	node.Tree = t
	return node
}

// ---------------------------- marshaling ----------------------------

// JSONL implements rwxrob/json.AsJSON.
func (s *Tree[T]) JSON() ([]byte, error) { return json.Marshal(s) }

// JSONL implements rwxrob/json.AsJSON.
func (s *Tree[T]) JSONL() ([]byte, error) {
	return json.MarshalIndent(s, "  ", "  ")
}

// String implements rwxrob/json.Stringer and fmt.Stringer.
func (s Tree[T]) String() string {
	byt, err := s.JSON()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// StringLong implements rwxrob/json.Stringer.
func (s Tree[T]) StringLong() string {
	byt, err := s.JSONL()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// String implements rwxrob/json.Printer.
func (s *Tree[T]) Print() { fmt.Println(s.String()) }

// PrintLong implements rwxrob/json.Printer.
func (s *Tree[T]) PrintLong() { fmt.Println(s.StringLong()) }

// Log implements rwxrob/json.Logger.
func (s Tree[T]) Log() { log.Print(s.String()) }

// LogLong implements rwxrob/json.Logger.
func (s Tree[T]) LogLong() { each.Log(to.Lines(s.StringLong())) }
