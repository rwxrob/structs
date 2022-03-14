package types

import (
	"encoding/json"
	"fmt"
	"log"
)

const UNKNOWN = 0

// Types contains integer types with a lookup index (which must always
// be kept in sync. The type 0 is always UNKNOWN. Structures that embed
// Types should always require a slice of strings be passed to their
// Init method which handles the creation of the Types.Map. Types should
// usually not be changed once initialized.
type Types struct {
	Names Names `json:",omitempty"`
	Map   Map   `json:",omitempty"`
}

// Set sets both the Names and Map to the new types ensuring that
// 0 (UNKNOWN) remains reserved.
func (t *Types) Set(types ...string) {
	t.Names = []string{"UNKNOWN"}
	t.Names = append(t.Names, types...)
	t.Map = map[string]int{"UNKNOWN": 0}
	for n, v := range types {
		t.Map[v] = n + 1
	}
}

// JSONL implements rwxrob/json.AsJSON.
func (s *Types) JSON() ([]byte, error) { return json.Marshal(s) }

// JSONL implements rwxrob/json.AsJSON.
func (s *Types) JSONL() ([]byte, error) {
	return json.MarshalIndent(s, "  ", "  ")
}

// String implements rwxrob/json.Stringer and fmt.Stringer.
func (s Types) String() string {
	byt, err := s.JSON()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// StringLong implements rwxrob/json.Stringer.
func (s Types) StringLong() string {
	byt, err := s.JSONL()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// String implements rwxrob/json.Printer.
func (s *Types) Print() { fmt.Println(s.String()) }

// PrintLong implements rwxrob/json.Printer.
func (s *Types) PrintLong() { fmt.Println(s.StringLong()) }

// Log implements rwxrob/json.Logger.
func (s Types) Log() { log.Print(s.String()) }

// LogLong implements rwxrob/json.Logger.
func (s Types) LogLong() { log.Print(s.StringLong()) }

// ---------------------------- Names ----------------------------

// Names allows names to be looked up by integer value. O is always
// UNKNOWN
type Names []string

// JSONL implements rwxrob/json.AsJSON.
func (s *Names) JSON() ([]byte, error) { return json.Marshal(s) }

// JSONL implements rwxrob/json.AsJSON.
func (s *Names) JSONL() ([]byte, error) {
	return json.MarshalIndent(s, "  ", "  ")
}

// String implements rwxrob/json.Stringer and fmt.Stringer.
func (s Names) String() string {
	byt, err := s.JSON()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// StringLong implements rwxrob/json.Stringer.
func (s Names) StringLong() string {
	byt, err := s.JSONL()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// String implements rwxrob/json.Printer.
func (s *Names) Print() { fmt.Println(s.String()) }

// PrintLong implements rwxrob/json.Printer.
func (s *Names) PrintLong() { fmt.Println(s.StringLong()) }

// Log implements rwxrob/json.Logger.
func (s Names) Log() { log.Print(s.String()) }

// LogLong implements rwxrob/json.Logger.
func (s Names) LogLong() { log.Print(s.StringLong()) }

// ----------------------------- Map -----------------------------

// Map allows type integer values to be looked up by name.
type Map map[string]int

// JSONL implements rwxrob/json.AsJSON.
func (s *Map) JSON() ([]byte, error) { return json.Marshal(s) }

// JSONL implements rwxrob/json.AsJSON.
func (s *Map) JSONL() ([]byte, error) {
	return json.MarshalIndent(s, "  ", "  ")
}

// String implements rwxrob/json.Stringer and fmt.Stringer.
func (s Map) String() string {
	byt, err := s.JSON()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// StringLong implements rwxrob/json.Stringer.
func (s Map) StringLong() string {
	byt, err := s.JSONL()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// String implements rwxrob/json.Printer.
func (s *Map) Print() { fmt.Println(s.String()) }

// PrintLong implements rwxrob/json.Printer.
func (s *Map) PrintLong() { fmt.Println(s.StringLong()) }

// Log implements rwxrob/json.Logger.
func (s Map) Log() { log.Print(s.String()) }

// LogLong implements rwxrob/json.Logger.
func (s Map) LogLong() { log.Print(s.StringLong()) }
