package functions

import ()

type Symbol interface {
	Symbol() // representation in source code
}

const (
	ForNumbers = "ForNumbers"
	ForStrings = "ForStrings"
)

type ElementType float64
type SliceType []ElementType

type Template string

var (
	templates = map[string][]string{
		"Sum.go":    []string{ForNumbers, ForStrings},
		"Filter.go": []string{ForNumbers, ForStrings},
		"Init.go":   []string{ForNumbers, ForStrings},
	}
)

func Templates() map[string][]string {
	return templates
}
