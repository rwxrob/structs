# Data Structs with Go (1.18+) Generics

[![Go
Version](https://img.shields.io/github/go-mod/go-version/rwxrob/structs)](https://tip.golang.org/doc/go1.18)
[![GoDoc](https://godoc.org/github.com/rwxrob/structs?status.svg)](https://godoc.org/github.com/rwxrob/structs)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report
Card](https://goreportcard.com/badge/github.com/rwxrob/structs)](https://goreportcard.com/report/github.com/rwxrob/structs)

* [Interfaces](interfaces.go)
* [Text Sets](set/text/set)
* [QStack](qstack_test.go)
* [Node](node_test.go)
* [Trees](tree)

All structures make judicious use of generics and implement the same
json.AsJSON interface (and others) making them much more consumable and
shareable.

## Design Decisions

* **Why no linked-list or queue?** Because they are fulfilled by QStack
  and Node.

* **Decided to drop async walks of Node.** It's easily accomplished by
  enclosing whatever channel is needed in the iterator function and
  forking a goroutine off from within it.

## TODO

* Add equivalent methods to Node from JavaScript (InsertAfter, etc.)
* Add Union and other set methods
