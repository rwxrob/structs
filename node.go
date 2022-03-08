package structs

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	UNKNOWN = 0
)

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
	T int      `json:"T"`          // type
	V T        `json:",omitempty"` // value
	P *Node[T] `json:"-"`          // up/parent

	// cached properties for efficient checks
	Count int // count of nodes under

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
	n.Count++
	u := new(Node[T])
	u.T = t
	u.V = v
	u.P = n
	if n.first == nil {
		n.first = u
		n.last = u
		return u
	}
	n.last.right = u
	u.left = n.last
	n.last = u
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
	from.first = nil
	from.last = nil
}

// ------------------------------- Walk --------------------------------

// WalkDeepPre will pass each Node in the tree to the given function
// traversing in a synchronous, depth-first, preorder way. The function
// passed may be a closure containing variables, contexts, or a channel
// outside of its own scope to be updated for each visit. This method
// uses functional recursion which may have some limitations depending
// on the depth of node trees required.
func (n *Node[T]) WalkDeepPre(do func(n *Node[T])) {
	list := []*Node[T]{n}
	for len(list) > 0 {
		cur := list[0]
		list = list[1:]
		nlist := []*Node[T]{} // TODO replace with QStack
		nlist = append(nlist, cur.Nodes()...)
		nlist = append(nlist, list...)
		list = nlist
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
