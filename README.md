# Data Structs with Go (1.18+) Generics

[![Go
Version](https://img.shields.io/github/go-mod/go-version/rwxrob/structs)](https://tip.golang.org/doc/go1.18)
[![GoDoc](https://godoc.org/github.com/rwxrob/structs?status.svg)](https://godoc.org/github.com/rwxrob/structs)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report
Card](https://goreportcard.com/badge/github.com/rwxrob/structs)](https://goreportcard.com/report/github.com/rwxrob/structs)

*Go report card has no idea how to handle Go 1.18.* 😀

* [Types Map](types)
* [QStack](qstack)
* [Rooted Node Tree](tree)
* [Text Sets](set/text/set)

All structures make judicious use of generics and implement the same
json.AsJSON interface (and others) making them much more consumable and
shareable.

## Design Decisions

* **Why no linked-list or queue?** Because they are fulfilled by QStack
  and Node.

* **Decided to drop async walks of Node.** It's easily accomplished by
  enclosing whatever channel is needed in the iterator function and
  forking a goroutine off from within it.

* **Don't stutter.** `tree.Tree` was changed to `tree.E` and
  `qstack.QStack` was changed to `qstack.QS` to follow the idiomatic
  "don't stutter" design preferred by core Go libraries like
  `testing.T`. (We don't want to be another `context.Context`.) 

## TODO

* Add equivalent methods to Node from JavaScript (InsertAfter, etc.)
* Add Union and other set methods
