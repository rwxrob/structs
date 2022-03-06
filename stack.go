package structs

// Stack specifies a basic stack data structure. Stacks can hold
// "stackables" of any type, including mixing different types in the
// same Stack. For a usable implementation see util.Stack but always use
// bonzai.Stack in functions signatures to maximize compatibility.
type Stack interface {
	Push(any)
	Pop() any
	Peek() any
}
