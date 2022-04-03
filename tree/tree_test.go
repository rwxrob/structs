package tree_test

import (
	"fmt"

	"github.com/rwxrob/structs/tree"
)

func ExampleTree() {
	t := tree.New[any]("foo")
	t.Print()
	t.Root.Print()
	fmt.Println(t.Types.Names)
	// Output:
	// {"Names":["UNKNOWN","foo"],"Map":{"UNKNOWN":0,"foo":1},"Root":{"T":1}}
	// {"T":1}
	// ["UNKNOWN","foo"]
}

func ExampleTree_Node() {
	t := tree.New[any]("foo")
	n := t.Node(10, "")
	n.Print()
	t.Print()
	t.Root.Append(n)
	t.Root.Print()
	// Output:
	// {"T":10,"V":""}
	// {"Names":["UNKNOWN","foo"],"Map":{"UNKNOWN":0,"foo":1},"Root":{"T":1}}
	// {"T":1,"N":[{"T":10,"V":""}]}
}
