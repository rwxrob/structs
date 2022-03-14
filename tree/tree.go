// Copyright 2022 Robert S. Muhlestein.
// SPDX-License-Identifier: Apache-2.0

package tree

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rwxrob/fn/each"
	"github.com/rwxrob/structs/qstack"
	"github.com/rwxrob/structs/types"
	"github.com/rwxrob/to"
)

// E ("tree-e") is an encapsulating struct to contain the Root Node and
// all possible Types for any Node. Most users of a tree will make
// direct use of E.Root (which can safely be swapped out for a root
// Node reference at any time). Tree implements the rwxrob/json/AsJSON
// interface and uses custom methods for MarshalJSON and UnmarshalJSON
// to facilitate storage, transfer, and documentation.
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

// Init initializes a tree creating its Root Node reference and assigning
// it the type of 1 (0 is reserved for UNKNOWN). Init then assigns its
// Types and indexes the TypesMap.
func (t *E[T]) Init(types []string) *E[T] {
	t.Types.Names = []string{"UNKNOWN"}
	t.Types.Names = append(t.Types.Names, types...)
	t.Types.Map = map[string]int{"UNKNOWN": 0}
	for n, v := range types {
		t.Types.Map[v] = n + 1
	}
	t.Root = new(Node[T])
	t.Root.T = 1
	t.Root.Tree = t
	return t
}

// Node returns new detached Node initialized for this same tree. Either
// an integer or string type may be passed and if none are passed the
// UNKNOWN (0) type will be assigned.
func (t *E[T]) Node(typ ...any) *Node[T] {
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

// LogLong implements rwxrob/json.Logger.
func (s E[T]) LogLong() { each.Log(to.Lines(s.StringLong())) }

// Node is an implementation of a "node" from traditional data
// structures. Nodes can have other nodes under then (which some call
// being a "parent" (P)). Nodes can also have a value (V). All nodes have
// a specific integer type (T) (which most will map to constants and
// slices of strings containing their readable name equivalent). The
// terms "up", "down", "under", "left", and "right" are consistent with
// the common visualization of a tree, but are not necessarily
// constraints on implementation or visual representation. Node is
// implemented as a linked list for maximum performance. Also see tree.
type Node[T any] struct {
	T     int      `json:"T"`          // type
	V     T        `json:",omitempty"` // value
	P     *Node[T] `json:"-"`          // up/parent
	Count int      `json:"-"`          // node count
	Tree  *E[T]    `json:"-"`          // optional tree with type names

	left  *Node[T]
	right *Node[T]
	first *Node[T]
	last  *Node[T]
}

// Init resets the node to its empty/zero state as if just created for
// the first time.
func (n *Node[T]) Init() {
	var zv T // required to get go's idea of zero value for instantiated type
	n.T = 0
	n.V = zv
	n.first = nil
	n.last = nil
	n.left = nil
	n.right = nil
}

// Nodes returns all the nodes under this Node. Prefer checking
// Count when the values are not needed.
func (n *Node[T]) Nodes() []*Node[T] {
	if n.first == nil {
		return nil
	}
	cur := n.first
	list := []*Node[T]{cur}
	for {
		cur = cur.right
		if cur == nil {
			break
		}
		list = append(list, cur)
	}
	return list
}

// --------------------------------------------------------------------

// Add creates a new Node with type and value under and returns. It also
// updates Count.
func (n *Node[T]) Add(t int, v T) *Node[T] {
	u := new(Node[T])
	u.T = t
	u.V = v
	u.P = n
	u.Tree = n.Tree
	n.Append(u)
	return u
}

// Cut removes a Node from under the one above it and returns.
func (n *Node[T]) Cut() *Node[T] {
	if n.left != nil {
		n.left.right = n.right
	}
	if n.right != nil {
		n.right.left = n.left
	}
	if n.P != nil {
		n.P.Count--
		if n == n.P.first {
			n.P.first = n.right
		}
		if n == n.P.last {
			n.P.last = n.left
		}
	}
	n.P = nil
	n.left = nil
	n.right = nil
	return n
}

// Take moves all nodes from another under itself.
func (n *Node[T]) Take(from *Node[T]) {
	if from.first == nil {
		return
	}
	if n.first == nil {
		n.first = from.first
		n.last = from.last
		n.Count = from.Count
	} else {
		n.Count += from.Count
		n.last.right = from.first
		from.first.left = n.last
		n.last = from.last
	}
	from.Count = 0
	from.first = nil
	from.last = nil
}

// Append adds an existing Node under this one as if Add had been
// called.
func (n *Node[T]) Append(u *Node[T]) {
	n.Count++
	if n.first == nil {
		n.first = u
		n.last = u
		return
	}
	n.last.right = u
	u.left = n.last
	n.last = u
}

// Morph initializes the node with Init and then sets it's value (V) and
// type (T) and all of its attachment references to those of the Node
// passed thereby preserving the Node reference of this method's
// receiver.
func (n *Node[T]) Morph(c *Node[T]) {
	n.Init()
	n.T = c.T
	n.V = c.V
	n.P = c.P
	n.left = c.left
	n.right = c.right
	n.first = c.first
	n.last = c.last
	n.Count = c.Count
}

// ------------------------------- Walk --------------------------------

// WalkLevels will pass each Node in the tree to the given function
// traversing in a synchronous, breadth-first, leveler way. The function
// passed may be a closure containing variables, contexts, or a channel
// outside of its own scope to be updated for each visit. This method
// uses functional recursion which may have some limitations depending
// on the depth of node trees required.
func (n *Node[T]) WalkLevels(do func(n *Node[T])) {
	list := qstack.New[*Node[T]]()
	list.Unshift(n)
	for list.Len > 0 {
		cur := list.Shift()
		list.Push(cur.Nodes()...)
		do(cur)
	}
}

// WalkDeepPre will pass each Node in the tree to the given function
// traversing in a synchronous, depth-first, preorder way. The function
// passed may be a closure containing variables, contexts, or a channel
// outside of its own scope to be updated for each visit. This method
// uses functional recursion which may have some limitations depending
// on the depth of node trees required.
func (n *Node[T]) WalkDeepPre(do func(n *Node[T])) {
	list := qstack.New[*Node[T]]()
	list.Unshift(n)
	for list.Len > 0 {
		cur := list.Shift()
		list.Unshift(cur.Nodes()...)
		do(cur)
	}
}

// ------------------------------ Printer -----------------------------

// just for marshaling
type jsnode[T any] struct {
	T int
	V T          `json:",omitempty"`
	N []*Node[T] `json:",omitempty"`
}

// MarshalJSON implements encoding/json.Marshaler and it needed to
// fulfill the list of nodes since they are internally stored as
// a linked list.
func (s Node[T]) MarshalJSON() ([]byte, error) {
	n := new(jsnode[T])
	n.T = s.T
	n.V = s.V
	n.N = s.Nodes()
	return json.Marshal(n)
}

// JSON implements rwxrob/json.Printer.
func (s *Node[T]) JSON() string {
	b, err := json.Marshal(s)
	if err != nil {
		log.Print(err)
	}
	return string(b)
}

// JSONL implements rwxrob/json.Printer.
func (s *Node[T]) JSONL() string {
	b, _ := json.MarshalIndent(s, "  ", "  ")
	return string(b)
}

// String implements rwxrob/json.Printer and fmt.Stringer.
func (s Node[T]) String() string { return s.JSON() }

// Print implements rwxrob/json.Printer.
func (s *Node[T]) Print() { fmt.Println(s.JSON()) }

// Log implements rwxrob/json.Printer.
func (s Node[T]) Log() { log.Print(s.JSON()) }

// PPrint implements rwxrob/json.Printer.
func (s *Node[T]) PPrint() { fmt.Println(s.JSONL()) }

// PLog implements rwxrob/json.Printer.
func (s Node[T]) LogL() { log.Print(s.JSONL()) }
