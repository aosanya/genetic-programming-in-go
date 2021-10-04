package main

import (
	"codevald.com/geneticprogramming"
)

func main() {
	b := geneticprogramming.Node{
		Function: geneticprogramming.FunctionWrapper{Name: "Add"},
	}
	b.Display()
}
