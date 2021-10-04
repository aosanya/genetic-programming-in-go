package geneticprogramming

type NodeFunction func(int) string

type FunctionWrapper struct {
	Name       string
	Function   NodeFunction
	ChildCount int
}
