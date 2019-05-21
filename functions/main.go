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
		"Abs.go": []string{ForNumbers},
		//		"Filter.go": []string{ForNumbers, ForStrings},
		//		"Head.go":   []string{ForNumbers, ForStrings},
		//		"Init.go":   []string{ForNumbers, ForStrings},
		//		"Last.go":   []string{ForNumbers, ForStrings},
		//		"Sum.go":    []string{ForNumbers, ForStrings},
		//		"Tail.go":   []string{ForNumbers, ForStrings},
	}
)

func Templates() map[string][]string {
	return templates
}
