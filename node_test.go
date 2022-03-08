package structs_test

import (
	"fmt"

	"github.com/rwxrob/structs"
)

func ExampleNode() {

	// flexible
	n := new(structs.Node[any])
	n.Print()

	// specific
	m := new(structs.Node[int])
	m.Print()

	// default type: UNKNOWN == 0
	fmt.Println(n.T == structs.UNKNOWN, m.T == structs.UNKNOWN)

	// values matching zero value are omitted

	// Output:
	// {"T":0}
	// {"T":0}
	// true true
}

func ExampleNode_type_is_Integer() {
	n := new(structs.Node[any])
	n.T = 1
	n.Print()
	//n.T = "one" // BOOM!

	// Output:
	// {"T":1}
}

func ExampleNode_value_Depends_on_Instantiation() {

	z := new(structs.Node[any])
	z.Print() // values that are equiv of zero value for type are omitted

	// flexible
	n := new(structs.Node[any])
	n.V = 1
	n.Print()
	n.V = true
	n.Print()
	n.V = "foo"
	n.Print()

	// strict
	m := new(structs.Node[int])
	m.V = 1
	m.Print()
	//m.V = true // BOOM!

	// Output:
	// {"T":0}
	// {"T":0,"V":1}
	// {"T":0,"V":true}
	// {"T":0,"V":"foo"}
	// {"T":0,"V":1}
}

func ExampleNode_Init() {

	// create and print a brand new one
	n := new(structs.Node[any])
	n.Print()

	// add something to it
	n.V = "something"
	n.Print()

	// initialize it back to "empty"
	n.Init()
	n.Print()

	// now try with something with a tricker zero value
	n.V = func() {}
	// log.Println(n.V) // "0x5019e0" yep it's there
	// n.Print()     // would log error and fail to marshal JSON

	// check it's properties
	fmt.Println(n.P != nil, n.Count)

	// Output:
	// {"T":0}
	// {"T":0,"V":"something"}
	// {"T":0}
	// false 0
}

func ExampleNode_properties() {

	// Nodes have these properties updating every time
	// their state is changed so that queries need not
	// to the checks again later.

	// initial state
	n := new(structs.Node[any])
	fmt.Println("n:", n.P == nil, n.V, n.Count)
	u := n.Add(1, nil)
	fmt.Println("n:", n.P == nil, n.V, n.Count)
	fmt.Println("u:", u.P == nil, u.V, u.Count)

	// make an edge node
	u.V = "something"

	// break edge by forcing it to have nodes and a value (discouraged)
	u.Add(9001, "muhaha")
	fmt.Println("u:", u.P == nil, u.V, u.Count)

	// Output:
	// n: true <nil> 0
	// n: true <nil> 1
	// u: false <nil> 0
	// u: false something 1

}

func ExampleNode_Nodes() {

	// create the first tree
	n := new(structs.Node[any])
	n.Add(1, nil)
	n.Add(2, nil)
	fmt.Println(n.Nodes(), n.Count)

	// and another added under it
	m := n.Add(3, nil)
	m.Add(3, nil)
	m.Add(3, nil)
	fmt.Println(m.Nodes(), m.Count)

	// Output:
	// [{"T":1} {"T":2}] 2
	// [{"T":3} {"T":3}] 2
}

func ExampleNode_Cut_middle() {
	n := new(structs.Node[any])
	n.Add(1, nil)
	c := n.Add(2, nil)
	n.Add(3, nil)
	n.Print()
	x := c.Cut()
	n.Print()
	x.Print()
	// Output:
	// {"T":0,"N":[{"T":1},{"T":2},{"T":3}]}
	// {"T":0,"N":[{"T":1},{"T":3}]}
	// {"T":2}
}

func ExampleNode_Cut_first() {
	n := new(structs.Node[any])
	c := n.Add(1, nil)
	n.Add(2, nil)
	n.Add(3, nil)
	n.Print()
	x := c.Cut()
	n.Print()
	x.Print()
	// Output:
	// {"T":0,"N":[{"T":1},{"T":2},{"T":3}]}
	// {"T":0,"N":[{"T":2},{"T":3}]}
	// {"T":1}
}

func ExampleNode_Cut_last() {
	n := new(structs.Node[any])
	n.Add(1, nil)
	n.Add(2, nil)
	c := n.Add(3, nil)
	n.Print()
	x := c.Cut()
	n.Print()
	x.Print()
	// Output:
	// {"T":0,"N":[{"T":1},{"T":2},{"T":3}]}
	// {"T":0,"N":[{"T":1},{"T":2}]}
	// {"T":3}
}

func ExampleNode_Take() {

	// build up the first
	n := new(structs.Node[any])
	n.T = 10
	n.Add(1, nil)
	n.Add(2, nil)
	n.Add(3, nil)
	n.Print()

	// now take them over

	m := new(structs.Node[any])
	m.T = 20
	m.Print()
	m.Take(n)
	m.Print()
	n.Print()

	// Output:
	// {"T":10,"N":[{"T":1},{"T":2},{"T":3}]}
	// {"T":20}
	// {"T":20,"N":[{"T":1},{"T":2},{"T":3}]}
	// {"T":10}

}

func ExampleNode_WalkLevels() {
	n := new(structs.Node[any])
	n.Add(1, nil).Add(11, nil)
	n.Add(2, nil).Add(22, nil)
	n.Add(3, nil).Add(33, nil)
	n.WalkLevels(func(c *structs.Node[any]) { fmt.Print(c.T, " ") })
	// Output:
	// 0 1 2 3 11 22 33
}

func ExampleNode_WalkDeepPre() {
	n := new(structs.Node[any])
	n.Add(1, nil).Add(11, nil)
	n.Add(2, nil).Add(22, nil)
	n.Add(3, nil).Add(33, nil)
	n.WalkDeepPre(func(c *structs.Node[any]) { fmt.Print(c.T, " ") })
	// Output:
	// 0 1 11 2 22 3 33
}

func ExampleNode_Morph() {
	n := new(structs.Node[any])
	n.Add(2, "some")
	m := new(structs.Node[any])
	m.Morph(n)
	n.Print()
	m.Print()
	// Output:
	// {"T":0,"N":[{"T":2,"V":"some"}]}
	// {"T":0,"N":[{"T":2,"V":"some"}]}
}
