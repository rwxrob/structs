package tree_test

import (
	"fmt"

	"github.com/rwxrob/structs/tree"
	"github.com/rwxrob/structs/types"
)

func ExampleNode() {

	n := new(tree.Node[any]) // any
	m := new(tree.Node[int]) // specific

	// default type: UNKNOWN == 0
	fmt.Println(n.T == types.UNKNOWN, m.T == types.UNKNOWN)

	// Output:
	// true true
}

func ExampleNode_Print_empty() {

	m := new(tree.Node[any])

	// Print indirectly calls MarshalJSON

	// [] because no value and UNKNOWN (0) type
	m.Print()

	// [] because empty JSON array and UNKNOWN (0) type
	m.V = []string{}
	m.Print()

	// [] because empty JSON object and UNKNOWN (0) type
	m.V = map[string]int{}
	m.Print()

	// [] because empty JSON string and UNKNOWN (0) type
	m.V = ""
	m.Print()

	// [] because empty JSON int is zero and UNKNOWN (0) type
	m.V = 0
	m.Print()

	// [] because empty JSON float is zero and UNKNOWN (0) type
	m.V = 0.000
	m.Print()

	// [] because empty JSON float is zero and UNKNOWN (0) type
	m.V = -0.000
	m.Print()

	// [] because empty JSON nil and UNKNOWN (0) type
	m.V = nil
	m.Print()

	// Output:
	// []
	// []
	// []
	// []
	// []
	// []
	// []
	// []
}
