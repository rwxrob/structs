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

// QS ("queue stack") implements a combination of the traditional queue
// and stack data structures using a linked list with Len field for high
// performance. Sometimes you want more than what standard slices
// provide, for example, when needing to prepend to a slice, which is
// very inefficient using standard Go slices.
type QS[T any] struct {
	Len int
	top *item[T]
	bot *item[T]
}

// New returns a newly initialized QS of the given type of items.
func New[T any]() *QS[T] { return new(QS[T]) }

// Items returns the items of the stack as a slice with newest items on
// the next. To apply an iterator function over the items consider
// using the rwxrob/fn package with Map/Filter/Reduce/Each functions.
func (s *QS[T]) Items() []T {
	items := []T{}
	for cur := s.bot; cur != nil; cur = cur.next {
		items = append(items, cur.V)
	}
	return items
}

// Peek (stack) returns the current top value of the stack. Prefer Len to
// check for emptiness.
func (s *QS[T]) Peek() T {
	var rv T
	if s.bot == nil {
		return rv
	}
	return s.bot.V
}

func (s *QS[T]) Push(these ...T) {
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
func (s *QS[T]) Pop() T {
	var rv T
	switch s.Len {
	case 0:
		return rv
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
func (s *QS[T]) Shift() T {
	var rv T
	switch s.Len {
	case 0:
		return rv
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
func (s *QS[T]) Unshift(these ...T) {
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
func (s QS[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Items())
}

// JSONL implements rwxrob/json.AsJSON.
func (s *QS[T]) JSON() ([]byte, error) { return s.MarshalJSON() }

// JSONL implements rwxrob/json.AsJSON.
func (s *QS[T]) JSONL() ([]byte, error) {
	return json.MarshalIndent(s, "  ", "  ")
}

// String implements rwxrob/json.Stringer and fmt.Stringer.
func (s QS[T]) String() string {
	byt, err := s.JSON()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// StringLong implements rwxrob/json.Stringer.
func (s QS[T]) StringLong() string {
	byt, err := s.JSONL()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// String implements rwxrob/json.Printer.
func (s *QS[T]) Print() { fmt.Println(s.String()) }

// PrintLong implements rwxrob/json.Printer.
func (s *QS[T]) PrintLong() { fmt.Println(s.StringLong()) }

// Log implements rwxrob/json.Logger.
func (s QS[T]) Log() { log.Print(s.String()) }

// LogLong implements rwxrob/json.Logger.
func (s QS[T]) LogLong() { log.Print(s.StringLong()) }

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
