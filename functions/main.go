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
		"abs.go":         {ForNumbers},
		"all.go":         {ForNumbers, ForStrings, ForStructs},
		"any.go":         {ForNumbers, ForStrings, ForStructs},
		"average.go":     {ForNumbers},
		"delete.go":      {ForNumbers, ForStrings, ForStructs},
		"filter.go":      {ForNumbers, ForStrings, ForStructs},
		"head.go":        {ForNumbers, ForStrings, ForStructs},
		"init.go":        {ForNumbers, ForStrings, ForStructs},
		"intercalate.go": {ForNumbers, ForStrings, ForStructs},
		"intersperse.go": {ForNumbers, ForStrings, ForStructs},
		"last.go":        {ForNumbers, ForStrings, ForStructs},
		"length.go":      {ForNumbers, ForStrings, ForStructs},
		"map.go":         {ForNumbers, ForStrings, ForStructs},
		"maximum.go":     {ForNumbers},
		"minimum.go":     {ForNumbers},
		"modes.go":       {ForNumbers, ForStrings, ForStructs},
		"nub.go":         {ForNumbers, ForStrings, ForStructs},
		"null.go":        {ForNumbers, ForStrings, ForStructs},
		"product.go":     {ForNumbers},
		"reverse.go":     {ForNumbers, ForStrings, ForStructs},
		"sort.go":        {ForNumbers, ForStrings},
		"sum.go":         {ForNumbers, ForStrings},
		"tail.go":        {ForNumbers, ForStrings, ForStructs},
		"take.go":        {ForNumbers, ForStrings, ForStructs},
		"uncons.go":      {ForNumbers, ForStrings, ForStructs},
		"unlines.go":     {ForNumbers, ForStrings, ForStructs},
		"unwords.go":     {ForNumbers, ForStrings, ForStructs},
	}
)

func Templates() map[string][]string {
	return templates
}
