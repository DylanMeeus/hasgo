package functions

import ()

type ElementType float64
type SliceType []ElementType

type Template string

var (
	templates = []string{"Sum.go"}
)

func Templates() []string {
	return templates
}
