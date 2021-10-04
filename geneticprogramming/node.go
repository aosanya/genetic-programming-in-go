package geneticprogramming

import (
	"fmt"
)

type Node struct {
	Function FunctionWrapper
}

func (a *Node) Name() string {
	return a.Function.Name
}

func (a *Node) Display() {
	fmt.Println(a.Name())
}
