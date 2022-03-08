package qstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type item[T any] struct {
	V    T
	prev *item[T]
	next *item[T]
}

// QStack implements a combination of the traditional queue and stack
// data structures using a linked list with Len field for high
// performance. Sometimes you want more than what standard slices
// provide, for example, when needing to prepend to a slice, which is
// very inefficient using standard Go slices.
type QStack[T any] struct {
	Len int
	top *item[T]
	bot *item[T]
}

// New returns a newly initialized QStack of the given type of items.
func New[T any]() *QStack[T] { return new(QStack[T]) }

// Items returns the items of the stack as a slice with newest items on
// the next. To apply an iterator function over the items consider
// using the rwxrob/fn package with Map/Filter/Reduce/Each functions.
func (s *QStack[T]) Items() []any {
	items := []any{}
	for cur := s.bot; cur != nil; cur = cur.next {
		items = append(items, cur.V)
	}
	return items
}

// Peek (stack) returns the current top value of the stack. Prefer Len to
// check for emptiness.
func (s *QStack[T]) Peek() any {
	if s.bot == nil {
		return nil
	}
	return s.bot.V
}

func (s *QStack[T]) Push(these ...T) {
	for i := 0; i < len(these); i++ {
		s.Len++
		n := new(item[T])
		n.V = these[i]
		switch s.Len {
		case 1:
			s.bot = n
			s.top = n
			continue
		default:
			s.top.next = n
			n.prev = s.top
			s.top = n
		}
	}
}

// Pop removes most recently pushed item from top of stack and returns it.
func (s *QStack[T]) Pop() any {
	switch s.Len {
	case 0:
		return nil
	case 1:
		s.Len--
		it := s.top
		s.top = nil
		return it.V
	default:
		s.Len--
		it := s.top
		s.top = s.top.prev
		if s.top != nil {
			s.top.next = nil
		}
		return it.V
	}
}

// Shift removes an item from the bottom of the stack and returns it.
func (s *QStack[T]) Shift() any {
	switch s.Len {
	case 0:
		return nil
	case 1:
		s.Len--
		it := s.bot
		s.bot = nil
		return it.V
	default:
		s.Len--
		it := s.bot
		s.bot = s.bot.next
		if s.bot != nil {
			s.bot.prev = nil
		}
		return it.V
	}
}

// Unshift adds items to the bottom of the stack.
func (s *QStack[T]) Unshift(these ...T) {
	for i := len(these) - 1; i >= 0; i-- {
		s.Len++
		n := new(item[T])
		n.V = these[i]
		switch s.Len {
		case 1:
			s.bot = n
			s.top = n
			continue
		default:
			s.bot.prev = n
			n.next = s.bot
			s.bot = n
		}
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

// ---------------------------- marshaling ----------------------------

// JSONL implements rwxrob/json.AsJSON.
func (s *item[T]) JSON() ([]byte, error) { return json.Marshal(s) }

// JSONL implements rwxrob/json.AsJSON.
func (s *item[T]) JSONL() ([]byte, error) {
	return json.MarshalIndent(s, "  ", "  ")
}

// String implements rwxrob/json.Stringer and fmt.Stringer.
func (s item[T]) String() string {
	byt, err := s.JSON()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// StringLong implements rwxrob/json.Stringer.
func (s item[T]) StringLong() string {
	byt, err := s.JSONL()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// String implements rwxrob/json.Printer.
func (s *item[T]) Print() { fmt.Println(s.String()) }

// PrintLong implements rwxrob/json.Printer.
func (s *item[T]) PrintLong() { fmt.Println(s.StringLong()) }

// Log implements rwxrob/json.Logger.
func (s item[T]) Log() { log.Print(s.String()) }

// LogLong implements rwxrob/json.Logger.
func (s item[T]) LogLong() { log.Print(s.StringLong()) }
