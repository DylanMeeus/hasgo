package functions

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
type SliceSliceType [][]ElementType

type Template string

var (
	templates = map[string][]string{
		"abs.go":         []string{ForNumbers},
		"all.go":         []string{ForNumbers, ForStrings, ForStructs},
		"any.go":         []string{ForNumbers, ForStrings, ForStructs},
		"filter.go":      []string{ForNumbers, ForStrings, ForStructs},
		"head.go":        []string{ForNumbers, ForStrings, ForStructs},
		"init.go":        []string{ForNumbers, ForStrings, ForStructs},
		"intercalate.go": []string{ForNumbers, ForStrings, ForStructs},
		"last.go":        []string{ForNumbers, ForStrings, ForStructs},
		"length.go":      []string{ForNumbers, ForStrings, ForStructs},
		"map.go":         []string{ForNumbers, ForStrings, ForStructs},
		"maximum.go":     []string{ForNumbers},
		"minimum.go":     []string{ForNumbers},
		"mode.go":        []string{ForNumbers, ForStrings, ForStructs},
		"null.go":        []string{ForNumbers, ForStrings, ForStructs},
		"product.go":     []string{ForNumbers},
		"reverse.go":     []string{ForNumbers, ForStrings, ForStructs},
		"sort.go":        []string{ForNumbers, ForStrings},
		"sum.go":         []string{ForNumbers, ForStrings},
		"tail.go":        []string{ForNumbers, ForStrings, ForStructs},
		"take.go":        []string{ForNumbers, ForStrings, ForStructs},
		"uncons.go":      []string{ForNumbers, ForStrings, ForStructs},
	}
)

func Templates() map[string][]string {
	return templates
}
