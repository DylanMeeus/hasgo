package functions

import ()

type Symbol interface {
	Symbol() // representation in source code
}

const (
	ForNumbers = "ForNumbers"
	ForStrings = "ForStrings"
	ForStructs = "ForStructs"
)

type ElementType float64
type SliceType []ElementType

type Template string

var (
	templates = map[string][]string{
		"Abs.go":     []string{ForNumbers},
		"Filter.go":  []string{ForNumbers, ForStrings, ForStructs},
		"Head.go":    []string{ForNumbers, ForStrings, ForStructs},
		"Init.go":    []string{ForNumbers, ForStrings, ForStructs},
		"Last.go":    []string{ForNumbers, ForStrings, ForStructs},
		"Map.go":     []string{ForNumbers, ForStrings, ForStructs},
		"Maximum.go": []string{ForNumbers},
		"Minimum.go": []string{ForNumbers},
		"Reverse.go": []string{ForNumbers, ForStrings, ForStructs},
		"Sort.go":    []string{ForNumbers, ForStrings},
		"Sum.go":     []string{ForNumbers, ForStrings},
		"Tail.go":    []string{ForNumbers, ForStrings, ForStructs},
		"Uncons.go":  []string{ForNumbers, ForStrings, ForStructs},
	}
)

func Templates() map[string][]string {
	return templates
}
