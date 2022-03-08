package structs_test

import (
	"fmt"
	"log"
	"os"

	"github.com/rwxrob/structs"
)

func ExampleQStack_Push() {
	s := new(structs.QStack)
	s.Print()
	s.Push("it")
	s.Print()
	// Output:
	// []
	// ["it"]
}

func ExampleQStack_Peek() {
	s := new(structs.QStack)
	s.Print()
	s.Push("it")
	fmt.Println(s.Peek())
	// Output:
	// []
	// it
}

func ExampleQStack_Pop() {
	s := new(structs.QStack)
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

func ExampleQStack_Items() {
	s := new(structs.QStack)
	s.Push(1)
	s.Push(true)
	s.Push("foo")
	s.Push(map[string]int{"ten": 10, "twenty": 20})
	s.Print()
	// Output:
	// [1,true,"foo",{"ten":10,"twenty":20}]
}

func ExampleQStack_Shift() {
	s := new(structs.QStack)
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

func ExampleQStack_Unshift() {
	s := new(structs.QStack)
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

func ExampleQStack_Has_Shift_Unshift() {
	s := new(structs.QStack)
	fmt.Println(s.Has, s.Len)
	s.Unshift("foo")
	fmt.Println(s.Has, s.Len)
	s.Shift()
	fmt.Println(s.Has, s.Len)
	// Output:
	// false 0
	// true 1
	// false 0
}

func ExampleQStack_Has_Push_Pop() {
	s := new(structs.QStack)
	fmt.Println(s.Has, s.Len)
	s.Push("foo")
	fmt.Println(s.Has, s.Len)
	s.Pop()
	fmt.Println(s.Has, s.Len)
	// Output:
	// false 0
	// true 1
	// false 0
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

	s := new(structs.QStack)
	s.Push(func() {})
	s.Print()

	// Output:
	// json: unsupported type: func()

}
