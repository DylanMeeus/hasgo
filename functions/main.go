package functions

// Symbol is represents a data token to be parsed.
type Symbol interface {
	Symbol()
}

// Types for which the function can be generated
const (
	ForNumbers = "ForNumbers"
	ForStrings = "ForStrings"
	ForStructs = "ForStructs"
)

// ElementType represents a single entry in a slice.
type ElementType float64

// SliceType represents a 1D slice.
type SliceType []ElementType

// SliceSliceType represents a 2D slice.
type SliceSliceType [][]ElementType

// Template is the content of a function file.
type Template string

var (
	templates = map[string][]string{
		"abs.go":         {ForNumbers},
		"all.go":         {ForNumbers, ForStrings, ForStructs},
		"any.go":         {ForNumbers, ForStrings, ForStructs},
		"average.go":     {ForNumbers},
		"delete.go":      {ForNumbers, ForStrings, ForStructs},
		"drop.go":        {ForNumbers, ForStrings, ForStructs},
		"elem.go":        {ForNumbers, ForStrings, ForStructs},
		"filter.go":      {ForNumbers, ForStrings, ForStructs},
		"foldl.go":       {ForNumbers, ForStrings, ForStructs},
		"foldl1.go":      {ForNumbers, ForStrings, ForStructs},
		"head.go":        {ForNumbers, ForStrings, ForStructs},
		"init.go":        {ForNumbers, ForStrings, ForStructs},
		"intercalate.go": {ForNumbers, ForStrings, ForStructs},
		"intersperse.go": {ForNumbers, ForStrings, ForStructs},
		"last.go":        {ForNumbers, ForStrings, ForStructs},
		"length.go":      {ForNumbers, ForStrings, ForStructs},
		"map.go":         {ForNumbers, ForStrings, ForStructs},
		"maximum.go":     {ForNumbers},
		"maximumby.go":   {ForNumbers, ForStrings, ForStructs},
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
		"takewhile.go":   {ForNumbers, ForStrings, ForStructs},
		"uncons.go":      {ForNumbers, ForStrings, ForStructs},
		"unlines.go":     {ForNumbers, ForStrings, ForStructs},
		"unwords.go":     {ForNumbers, ForStrings, ForStructs},
	}
)

// Templates returns all entries in the template map
func Templates() map[string][]string {
	return templates
}
