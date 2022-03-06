package stack_test

import (
	"github.com/rwxrob/structs/stack"
)

func ExampleStack_New() {
	s := stack.New()

	// Output:
	// another
	// another
	// some
	// 1
	// <nil>
}
