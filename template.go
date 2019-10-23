// Code generated by go generate; DO NOT EDIT.
package main

var hasgoTemplates = map[string]string{
	"abs.go": `
import (
	"math"
)

// Abs returns the absolute value of all elements in the slice.
// Can be generated for number-types.
func (s SliceType) Abs() (out SliceType) {
	for _, v := range s {
		out = append(out, ElementType(math.Abs(float64(v))))
	}
	return
}
`,
	"all.go": `
// All returns true if all elements of the slice satisfy the predicate.
// Can be generated for any type.
func (s SliceType) All(f func(ElementType) bool) bool {
	if f == nil {
		return false
	}
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}
`,
	"any.go": `
// Any returns true if any of the elements satisfy the predicate.
// Can be generated for any type.
func (s SliceType) Any(f func(ElementType) bool) bool {
	if f == nil {
		return false
	}
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}
`,
	"average.go": `
// Average returns the average of all elements in the slice.
// Can be generated for all number types.
func (s SliceType) Average() float64 {
	var sum ElementType
	if len(s) == 0 {
		return float64(0)
	}
	for _, i := range s {
		sum += i
	}
	return float64(sum) / float64(len(s))
}
`,
	"delete.go": `
// Delete returns a slice with the first matching element
// removed from the slice.
// Can be generated for any type.
func (s SliceType) Delete(e ElementType) (out SliceType) {
	deleted := false
	for _, v := range s {
		if deleted || v != e {
			out = append(out, v)
		} else {
			deleted = true
		}
	}
	return
}
`,
	"drop.go": `
// Drop returns a new Slice with the elements after the provided key.
// Can be generated for any type.
func (s SliceType) Drop(i int) (out SliceType) {
	for i < len(s) {
		out = append(out, s[i])
		i++
	}
	return
}
`,
	"elem.go": `
// Elem returns true if the slice contains the element
// Can be generated for any type.
func (s SliceType) Elem(el ElementType) bool {
	for _, e := range s {
		if e == el {
			return true
		}
	}
	return false
}
`,
	"filter.go": `
// Filter returns a slice containing only the elements that match the predicate.
// Can be generated for any type.
func (s SliceType) Filter(f func(ElementType) bool) (out SliceType) {
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}
`,
	"foldl.go": `
// Foldl reduces a list by iteratively applying f from left->right. Thus, for an empty slice, the result is the default zero-value.
func (s SliceType) Foldl(z ElementType, f func(e1, e2 ElementType) ElementType) (out ElementType) {
	if len(s) == 0 {
		return
	}
	out = s[0]
	for _, v := range s[1:] {
		out = f(out, v)
	}
	return f(out, z)
}
`,
	"foldl1.go": `
// Foldl1 reduces a list by iteratively applying f from left->right. Thus, for an empty slice, the result is the default zero-value.
func (s SliceType) Foldl1(f func(e1, e2 ElementType) ElementType) (out ElementType) {
	if len(s) == 0 {
		return
	}
	out = s[0]
	for _, v := range s[1:] {
		out = f(out, v)
	}
	return
}
`,
	"head.go": `
// Head returns the first element in the slice.
// If no element is found, returns the zero-value of the type.
// Can be generated for any type.
func (s SliceType) Head() (out ElementType) {
	if len(s) > 0 {
		out = s[0]
	}
	return
}
`,
	"init.go": `
// Init takes n-1 elements from a slice, where n = len(list).
// Can be generated for any type.
func (s SliceType) Init() (out SliceType) {
	if len(s) == 0 {
		return
	}
	slicecopy := append([]ElementType(nil), s...)
	return slicecopy[:len(s)-1]
}
`,
	"intercalate.go": `
// Intercalate inserts the method receiver slice into the function slice at each step.
// Can be generated for any type.
func (s SliceType) Intercalate(ss SliceSliceType) (out SliceType) {
	for i, slice := range ss {
		for _, el := range slice {
			out = append(out, el)
		}
		if i == len(ss)-1 {
			break
		}
		for _, el := range s {
			out = append(out, el)
		}
	}
	return out
}
`,
	"intersperse.go": `
// Intersperse inserts the receiving value between each element of the method receiver.
// Can be generated for any type.
func (s SliceType) Intersperse(value ElementType) (out SliceType) {
	for i, el := range s {
		out = append(out, el)
		if i == len(s)-1 {
			break
		}
		out = append(out, value)
	}
	return
}
`,
	"last.go": `
// Last returns the last element in the slice
// If no element is found, returns the zero-value of the type
// Can be generated for any type.
func (s SliceType) Last() (out ElementType) {
	if len(s) > 0 {
		out = s[len(s)-1]
	}
	return
}
`,
	"length.go": `
// Length returns the length (len) of a slice.
// Can be generated for any type.
func (s SliceType) Length() int {
	return len(s)
}
`,
	"map.go": `
// Map return a new slice with the map operation applied to each element.
// Can be generated for any type.
func (s SliceType) Map(f func(ElementType) ElementType) (out SliceType) {
	if f == nil {
		return s
	}
	for _, v := range s {
		out = append(out, f(v))
	}
	return
}
`,
	"maximum.go": `
// Maximum returns the maximum in a slice.
// Can be generated for number types.
func (s SliceType) Maximum() (out ElementType) {
	if len(s) == 0 {
		return
	}
	for _, i := range s {
		if i > out {
			out = i
		}
	}
	return
}
`,
	"maximumby.go": `
// MaximumBy returns the maximum elements according to a custom comparator.
// Can be generated for any type.
func (s SliceType) MaximumBy(f func(e1, e2 ElementType) ElementType) (max ElementType) {
	if len(s) == 0 {
		return
	}
	max = s[0]
	for _, el := range s[1:] {
		max = f(max, el)
	}
	return
}
`,
	"minimum.go": `
// Minimum returns the minimum of a slice.
// Can be generated for number types.
func (s SliceType) Minimum() ElementType {
	if len(s) == 0 {
		return 0 // bit strange?
	}
	min := s[0]
	for _, i := range s {
		if i < min {
			min = i
		}
	}
	return min
}
`,
	"modes.go": `
// Modes returns the elements with highest frequency in the slice.
// Can be generated for any type.
func (s SliceType) Modes() (out SliceType) {
	if len(s) == 0 {
		return
	}

	counts := make(map[ElementType]int)
	for _, v := range s {
		counts[v]++
	}

	var max int
	for _, v := range counts {
		if v > max {
			max = v
		}
	}

	for k, v := range counts {
		if v == max {
			out = append(out, k)
		}
	}

	return
}
`,
	"nub.go": `
// Nub returns a slice containing only the unique elements of the receiver.
// The order of the elements is preserved.
// Can be generated for any type.
func (s SliceType) Nub() (out SliceType) {
	if len(s) == 0 {
		return
	}

	contains := make(map[ElementType]struct{})
	for _, v := range s {
		if _, ok := contains[v]; !ok {
			contains[v] = struct{}{}
			out = append(out, v)
		}
	}
	return
}
`,
	"null.go": `
// Null returns true the slice is empty.
// Can be generated for any type.
func (s SliceType) Null() bool {
	return len(s) == 0
}
`,
	"product.go": `
// Product returns the product of all elements in the slice.
// Can be generated for any number type.
func (s SliceType) Product() ElementType {
	var prod ElementType
	for _, v := range s {
		prod += v
	}
	return prod
}
`,
	"reverse.go": `
// Reverse returns the reversed slice.
// Can be generated for any type.
func (s SliceType) Reverse() (out SliceType) {
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return
}
`,
	"sort.go": `
import (
	"sort"
)

// Sort is a wrapper around go sort function.
// Can be generated for any type.
func (s SliceType) Sort() SliceType {
	out := make(SliceType, len(s))
	copy(out, s)
	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})
	return out
}
`,
	"sum.go": `
// Sum returns the sum of all elements in the slice.
// Can be generated for any number type.
func (s SliceType) Sum() ElementType {
	var sum ElementType
	for _, v := range s {
		sum += v
	}
	return sum
}
`,
	"tail.go": `
// Tail takes [1 -> n] elements from a slice, where n = len(list)
// Returns an empty slice if there are less than 2 elements in slice
// Can be generated for any type.
func (s SliceType) Tail() (out SliceType) {
	if len(s) <= 1 {
		return
	}
	slicecopy := append([]ElementType(nil), s...)
	return slicecopy[1:]
}
`,
	"take.go": `
// Take takes the first n elements of the slice, or the entire slice if n > len(slice).
// Can be generated for any type.
func (s SliceType) Take(n uint64) (out SliceType) {
	if len(s) == 0 {
		return
	}
	out = make(SliceType, len(s))
	copy(out, s)
	if n < uint64(len(s)) {
		return out[:n]
	}
	return
}
`,
	"takewhile.go": `
// TakeWhile continues appending to the output as long as the predicate is satisfied.
// Can be generated for any type.
func (s SliceType) TakeWhile(p func(ElementType) bool) (out SliceType) {
	if len(s) == 0 {
		return
	}
	for _, e := range s {
		if !p(e) {
			return
		}
		out = append(out, e)
	}
	return
}
`,
	"uncons.go": `
// Uncons decomposes a slice into the head and tail component.
// Can be generated for any type.
func (s SliceType) Uncons() (head ElementType, tail SliceType) {
	return s.Head(), s.Tail()
}
`,
	"unlines.go": `
import (
	"fmt"
)

// Unlines joins together the string representation of the slice with newlines after each element.
// Can be generated for any type.
func (s SliceType) Unlines() (out string) {
	for i, v := range s {
		out += fmt.Sprintf("%v", v)
		if i != len(s)-1 {
			out += "\n"
		}
	}
	return
}
`,
	"unwords.go": `
import (
	"fmt"
)

// Unwords joins together the string representation of the slice with newlines after each element.
// Can be generated for any type.
func (s SliceType) Unwords() (out string) {
	for i, v := range s {
		out += fmt.Sprintf("%v", v)
		if i != len(s)-1 {
			out += " "
		}
	}
	return
}
`,
}

const (
	ForNumbers = "ForNumbers"
	ForStrings = "ForStrings"
	ForStructs = "ForStructs"
)

var funcDomains = map[string][]string{
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
