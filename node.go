package structs

// Node specifies a node (or leaf/edge) for use specifically in a rooted
// node tree data structure (the most common and practical). The Nodes
// under a Node are usually implemented as a linked list and therefore
// checking HasNodes first before calling Nodes is usually preferred.
// Also see bonzai.Tree interface.
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
	Detach()          // detach from node this node is under
	Take(from Node)   // take all nodes from under another
}
