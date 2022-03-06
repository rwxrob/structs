package structs

import "context"

// Stack specifies a basic stack data structure. Stacks can hold
// "stackables" of any type, including mixing different types in the
// same Stack. For a usable implementation see the stack subpackage.
type Stack interface {
	Push(any)
	Pop() any
	Peek() any
}

// Tree specifies a rooted node tree made up of Node implementations.
// The Tree holds the meta data about Node types including their string
// equivalents corresponding to integer type. Implementations of Tree
// must also implement Node fully as well so that the underlying structs
// work well together, specifically the Node struct implementation
// should know about the Tree struct implementation to which it belongs.
// When a Tree is initialized with Init it sets the internal cache of
// type strings and (re)generates the Root node. All other Nodes are
// added to the tree with Node.Add. The Tree.Root Node can never be
// destroyed (Node.Destroy) since Tree.Root will always refer to it even
// if it has no Nodes of its own and an empty value.
type Tree interface {
	Init(ts []string) error  // (re)initialize with new types and new root
	Types() []string         // returns slice corresponding to type int
	TypeString(i int) string // returns string for given type
	Node(t int) Node         // return new node of type t for tree
	Root() Node              // returns root node
}

// Node specifies a node (or leaf/edge) for use in any graph or tree
// data structure. Nodes under a Node are usually implemented as
// a linked list and therefore checking HasNodes first before calling
// Nodes is usually preferred. See the node subpackage for a usable
// implementation.
type Node interface {
	Type() int        // constants with string names
	SetType(int)      // possible, but usually avoid
	Value() any       // usually "empty" if edge
	Node() Node       // node that this node is under
	Nodes() []Node    // nodes under this node
	HasNodes() bool   // has nodes under it, check first
	IsRoot() bool     // not under another but has under self
	IsEdge() bool     // nothing under but has a value
	IsEmpty() bool    // nothing under and no value
	Init(t int) error // set type and any state for that type
	Add(t int) Node   // add a new node under self
	Destroy()         // removes self from existence
	Cut()             // detach from upper node but not under
	Take(from Node)   // take all nodes from under another
}

// VisitsAsync specifies a collection type that implements a Visit
// method that will perform the given function on all items in the
// collection sending their returned values to the given channel.
// A limit may be placed on the number of concurrent visit functions
// that may run at the same time. A limit of 0 indicates no limit. The
// method of visiting items is unspecified. The returned error may be
// used either to indicate errors with the setup and execution of the
// visit or to indicate the overall success of the entire visit. The
// context allows concurrent visits to be safely cancelled
// simultaneously (timeouts, etc.). Also see Visits. For specific
// example of struct that implements see the node subpackage.
type VisitsAsync[T any] interface {
	Visit(c context.Context, f func(item T) any, lim int, rv chan any) error
}

// Visits specifies a collection type that implements a Visit method
// that will perform the given function on all items in the collection
// (synchronously) sending their returned values to the given channel as
// each completes. The method of visiting items is unspecified other than
// ensuring every item is visited synchronously (func completes before
// another begins). The returned error may be used either to indicate
// errors with the setup and execution of the visit or to indicate the
// overall success of the entire visit.  Also see VisitsAsync. For
// a specific example of struct that implements see the node subpackage.
type Visits[T any] interface {
	Visit(f func(item T) any, rv chan any) error
}
