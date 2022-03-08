package structs

import (
	"encoding/json"
	"fmt"
	"log"
)

type item[T any] struct {
	val   T
	left  *item[T]
	right *item[T]
}

// QStack implements a combination of the traditional queue and stack
// data structures using a linked list with Has and Len fields
// for high performance. Sometimes you want more than what standard
// slices provide.
type QStack[T any] struct {
	Has bool
	Len int
	top *item[T]
	bot *item[T]
}

// Items returns the items of the stack as a slice with newest items on
// the right. To apply an iterator function over the items consider
// using the rwxrob/fn package with Map/Filter/Reduce/Each functions.
func (s *QStack[T]) Items() []any {
	items := []any{}
	for cur := s.top; cur != nil; cur = cur.right {
		items = append(items, cur.val)
	}
	return items
}

// Peek (stack) returns the current top value of the stack. Do not
// use to check if empty or not (use Has instead).
func (s *QStack[T]) Peek() any {
	if s.bot == nil {
		return nil
	}
	return s.bot.val
}

// Push (stack) will add anything to the top.
func (s *QStack[T]) Push(these ...T) {
	for i := len(these) - 1; i >= 0; i-- {
		s.Has = true
		s.Len++
		n := new(item[T])
		n.val = these[i]
		if s.top == nil {
			s.top = n
			s.bot = n
			return
		}
		s.bot.right = n
		n.left = s.bot
		s.bot = n
	}
}

// Pop removes most recently pushed item from top of stack and returns it.
func (s *QStack[T]) Pop() any {
	if s.bot == nil {
		return nil
	}
	popped := s.bot
	s.bot = s.bot.left
	if s.bot == nil {
		s.top = nil
		s.Has = false
		s.Len = 0
		return popped.val
	}
	s.bot.right = nil
	return popped.val
}

// Shift removes an item from the bottom of the stack and returns it.
func (s *QStack[T]) Shift() any {
	if s.top == nil {
		return nil
	}
	shifted := s.top
	s.top = s.top.right
	if s.top == nil {
		s.top = nil
		s.Has = false
		s.Len = 0
		return shifted.val
	}
	s.top.left = nil
	return shifted.val
}

// Unshift adds items to the bottom of the stack.
func (s *QStack[T]) Unshift(these ...T) {
	for i := len(these) - 1; i >= 0; i-- {
		s.Has = true
		s.Len++
		n := new(item[T])
		n.val = these[i]
		if s.top == nil {
			s.top = n
			s.bot = n
			return
		}
		s.top.left = n
		n.right = s.top
		s.top = n
	}
}

// ---------------------------- marshaling ----------------------------

// MarshalJSON implements encoding/json.Marshaler and it needed to
// fulfill the list of nodes since they are internally stored as
// a linked list.
func (s QStack[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Items())
}

// JSONL implements rwxrob/json.AsJSON.
func (s *QStack[T]) JSON() ([]byte, error) { return s.MarshalJSON() }

// JSONL implements rwxrob/json.AsJSON.
func (s *QStack[T]) JSONL() ([]byte, error) {
	return json.MarshalIndent(s, "  ", "  ")
}

// String implements rwxrob/json.Stringer and fmt.Stringer.
func (s QStack[T]) String() string {
	byt, err := s.JSON()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// StringLong implements rwxrob/json.Stringer.
func (s QStack[T]) StringLong() string {
	byt, err := s.JSONL()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// String implements rwxrob/json.Printer.
func (s *QStack[T]) Print() { fmt.Println(s.String()) }

// PrintLong implements rwxrob/json.Printer.
func (s *QStack[T]) PrintLong() { fmt.Println(s.StringLong()) }

// Log implements rwxrob/json.Logger.
func (s QStack[T]) Log() { log.Print(s.String()) }

// LogLong implements rwxrob/json.Logger.
func (s QStack[T]) LogLong() { log.Print(s.StringLong()) }
