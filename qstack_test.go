package structs_test

import (
	"fmt"
	"log"
	"os"

	"github.com/rwxrob/structs"
)

func ExampleQStack_Push() {
	s := new(structs.QStack[any])
	s.Print()
	s.Push("one")
	s.Print()
	s.Push("two")
	s.Print()
	// Output:
	// []
	// ["one"]
	// ["one","two"]
}

func ExampleQStack_Peek() {
	s := new(structs.QStack[any])
	s.Print()
	s.Push("it")
	fmt.Println(s.Peek())
	// Output:
	// []
	// it
}

func ExampleQStack_Pop() {
	s := new(structs.QStack[any])
	s.Print()
	p := s.Pop()
	fmt.Println(p)
	s.Push("it")
	s.Push("again")
	s.Print()
	fmt.Println(s.Len)
	p = s.Pop()
	s.Print()
	fmt.Println(p)
	fmt.Println(s.Len)
	// Output:
	// []
	// <nil>
	// ["it","again"]
	// 2
	// ["it"]
	// again
	// 1
}

func ExampleQStack_Items() {
	s := new(structs.QStack[any])
	s.Push(1)
	s.Push(true)
	s.Push("foo")
	s.Push(map[string]int{"ten": 10, "twenty": 20})
	s.Print()
	// Output:
	// [1,true,"foo",{"ten":10,"twenty":20}]
}

func ExampleQStack_Shift() {
	s := new(structs.QStack[any])
	s.Push(1)
	s.Push(true)
	s.Push("foo")
	v := s.Shift()
	fmt.Println(v, s.Len)
	s.Print()
	v = s.Shift()
	fmt.Println(v, s.Len)
	s.Print()
	v = s.Shift()
	fmt.Println(v, s.Len)
	s.Print()
	// Output:
	// 1 2
	// [true,"foo"]
	// true 1
	// ["foo"]
	// foo 0
	// []
}

func ExampleQStack_Unshift() {
	s := new(structs.QStack[any])
	s.Push(1)
	s.Push(true)
	s.Push("foo")
	s.Print()
	fmt.Println(s.Len)
	s.Unshift(0, 34, 2)
	s.Print()
	fmt.Println(s.Len)
	// Output:
	// [1,true,"foo"]
	// 3
	// [0,34,2,1,true,"foo"]
	// 6
}

func ExampleQStack_invalid_JSON_Types() {
	defer log.SetOutput(os.Stderr)
	defer log.SetFlags(log.Flags())
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	// QStack can be used to store any type,
	// but log an error (no panic) when
	// attempting to use the stack item
	// in a string context.

	s := new(structs.QStack[any])
	s.Push(func() {})
	s.Print()

	// Output:
	// json: unsupported type: func()

}
