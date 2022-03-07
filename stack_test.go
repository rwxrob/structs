package structs_test

import (
	"fmt"
	"log"
	"os"

	"github.com/rwxrob/structs"
)

func ExampleStack_Push() {
	s := new(structs.Stack)
	s.Print()
	s.Push("it")
	s.Print()
	// Output:
	// []
	// ["it"]
}

func ExampleStack_Peek() {
	s := new(structs.Stack)
	s.Print()
	s.Push("it")
	fmt.Println(s.Peek())
	// Output:
	// []
	// it
}

func ExampleStack_Pop() {
	s := new(structs.Stack)
	s.Print()
	p := s.Pop()
	fmt.Println(p)
	s.Push("it")
	s.Print()
	p = s.Pop()
	s.Print()
	fmt.Println(p)
	// Output:
	// []
	// <nil>
	// ["it"]
	// []
	// it
}

func ExampleStack_Items() {
	s := new(structs.Stack)
	s.Push(1)
	s.Push(true)
	s.Push("foo")
	s.Push(map[string]int{"ten": 10, "twenty": 20})
	s.Print()
	// Output:
	// [1,true,"foo",{"ten":10,"twenty":20}]
}

func ExampleStack_Shift() {
	s := new(structs.Stack)
	s.Push(1)
	s.Push(true)
	s.Push("foo")
	v := s.Shift()
	fmt.Println(v)
	s.Print()
	// Output:
	// 1
	// [true,"foo"]
}

func ExampleStack_Unshift() {
	s := new(structs.Stack)
	s.Push(1)
	s.Push(true)
	s.Push("foo")
	s.Print()
	s.Unshift(0)
	s.Print()
	// Output:
	// [1,true,"foo"]
	// [0,1,true,"foo"]
}

func ExampleStack_Has_Shift_Unshift() {
	s := new(structs.Stack)
	fmt.Println(s.Has)
	s.Unshift("foo")
	fmt.Println(s.Has)
	s.Shift()
	fmt.Println(s.Has)
	// Output:
	// false
	// true
	// false
}

func ExampleStack_Has_Push_Pop() {
	s := new(structs.Stack)
	fmt.Println(s.Has)
	s.Push("foo")
	fmt.Println(s.Has)
	s.Pop()
	fmt.Println(s.Has)
	// Output:
	// false
	// true
	// false
}

func ExampleStack_invalid_JSON_Types() {
	defer log.SetOutput(os.Stderr)
	defer log.SetFlags(log.Flags())
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	// Stack can be used to store any type,
	// but log an error (no panic) when
	// attempting to use the stack item
	// in a string context.

	s := new(structs.Stack)
	s.Push(func() {})
	s.Print()

	// Output:
	// json: unsupported type: func()

}
