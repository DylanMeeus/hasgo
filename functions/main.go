package functions

import ()

type Symbol interface {
	Symbol() // representation in source code
}

type ElementType float64
type SliceType []ElementType

type Template string

var (
	templates = []string{"Sum.go"}
)

func Templates() []string {
	return templates
}
