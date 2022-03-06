package structs

import (
	"context"
)

// AsyncVisitor specifies a collection type that implements a Visit
// method that will perform the given function on all items in the
// collection sending their returned values to the given channel.
// A limit may be placed on the number of concurrent visit functions
// that may run at the same time. A limit of 0 indicates no limit. The
// method of visiting items is unspecified. The returned error may be
// used either to indicate errors with the setup and execution of the
// visit or to indicate the overall success of the entire visit. The
// context allows concurrent visits to be safely cancelled
// simultaneously (timeouts, etc.). Also see Visitable.
type VisitsAsync[T any] interface {
	Visit(c context.Context, f func(item T) any, lim int, rv chan any) error
}

// Visits specifies a collection type that implements a Visit method
// that will perform the given function on all items in the collection
// (synchronously) sending their returned values to the given channel as
// each completes. The method of visiting items is unspecified other than
// ensuring every item is visited synchronously (func completes before
// another begins). The returned error may be used either to indicate
// errors with the setup and execution of the visit or to indicate the
// overall success of the entire visit.  Also see VisitsAsync.
type Visits[T any] interface {
	Visit(f func(item T) any, rv chan any) error
}
