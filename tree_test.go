package structs_test

import (
	"fmt"

	"github.com/rwxrob/structs"
)

func ExampleTree() {
	t := new(structs.Tree[any])
	t.Print()
	t.Init([]string{"main", "foo", "bar"})
	t.Root.Print()
	fmt.Println(t.Types)
	// Output:
	// {}
	// {"T":1}
	// [UNKNOWN main foo bar]
}

func ExampleTree_Node() {
	t := new(structs.Tree[any])
	t.Init([]string{"foo"})
	n := t.Node()
	n.Print()
	t.Print()
	t.Root.Append(n)
	t.Root.Print()
	// Output:
	// {"T":0}
	// {"Root":{"T":1},"Types":["UNKNOWN","foo"],"TypesMap":{"UNKNOWN":0,"foo":1}}
	// {"T":1,"N":[{"T":0}]}
}
