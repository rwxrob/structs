package structs

import (
	"encoding/json"
	"fmt"
	"log"
)

type stacked struct {
	val   any
	left  *stacked
	right *stacked
}

// Stack implements a performant stack data structure using a linked list.
type Stack struct {
	Has   bool
	first *stacked
	last  *stacked
}

// Items returns the items of the stack as a slice with newest items on
// the right.
func (s *Stack) Items() []any {
	items := []any{}
	for cur := s.first; cur != nil; cur = cur.right {
		items = append(items, cur.val)
	}
	return items
}

// Peek returns the current top value of the stack. Do not use to check if
// empty or not (use Has instead).
func (s *Stack) Peek() any {
	if s.last == nil {
		return nil
	}
	return s.last.val
}

// Push will add an item of any type on top of the stack.
func (s *Stack) Push(it any) {
	s.Has = true
	n := new(stacked)
	n.val = it
	if s.first == nil {
		s.first = n
		s.last = n
		return
	}
	s.last.right = n
	n.left = s.last
	s.last = n
}

// Pop removes most recently pushed item from top of stack and returns it.
func (s *Stack) Pop() any {
	if s.last == nil {
		return nil
	}
	popped := s.last
	s.last = s.last.left
	if s.last == nil {
		s.first = nil
		s.Has = false
		return popped.val
	}
	s.last.right = nil
	return popped.val
}

// Shift removes an item from the bottom of the stack and returns it.
func (s *Stack) Shift() any {
	if s.first == nil {
		return nil
	}
	shifted := s.first
	s.first = s.first.right
	if s.first == nil {
		s.first = nil
		s.Has = false
		return shifted.val
	}
	s.first.left = nil
	return shifted.val
}

// Unshift adds an item to the bottom of the stack.
func (s *Stack) Unshift(it any) {
	s.Has = true
	n := new(stacked)
	n.val = it
	if s.first == nil {
		s.first = n
		s.last = n
		return
	}
	s.first.left = n
	n.right = s.first
	s.first = n
}

// ---------------------------- marshaling ----------------------------

// MarshalJSON implements encoding/json.Marshaler and it needed to
// fulfill the list of nodes since they are internally stored as
// a linked list.
func (s Stack) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Items())
}

// JSONL implements rwxrob/json.AsJSON.
func (s *Stack) JSON() ([]byte, error) { return s.MarshalJSON() }

// JSONL implements rwxrob/json.AsJSON.
func (s *Stack) JSONL() ([]byte, error) {
	return json.MarshalIndent(s, "  ", "  ")
}

// String implements rwxrob/json.Stringer and fmt.Stringer.
func (s Stack) String() string {
	byt, err := s.JSON()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// StringLong implements rwxrob/json.Stringer.
func (s Stack) StringLong() string {
	byt, err := s.JSONL()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// String implements rwxrob/json.Printer.
func (s *Stack) Print() { fmt.Println(s.String()) }

// PrintLong implements rwxrob/json.Printer.
func (s *Stack) PrintLong() { fmt.Println(s.StringLong()) }

// Log implements rwxrob/json.Logger.
func (s Stack) Log() { log.Print(s.String()) }

// LogLong implements rwxrob/json.Logger.
func (s Stack) LogLong() { log.Print(s.StringLong()) }
