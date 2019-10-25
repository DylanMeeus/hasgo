// code generated by hasgo. [DO NOT EDIT!]
package types

import (
	"fmt"
	"math"
	"sort"
)

// =============== abs.go =================

// Abs returns the absolute value of all elements in the slice.
// Can be generated for number-types.
func (s Ints) Abs() (out Ints) {
	for _, v := range s {
		out = append(out, int64(math.Abs(float64(v))))
	}
	return
}

// =============== all.go =================

// All returns true if all elements of the slice satisfy the predicate.
// Can be generated for any type.
func (s Ints) All(f func(int64) bool) bool {
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

// =============== any.go =================

// Any returns true if any of the elements satisfy the predicate.
// Can be generated for any type.
func (s Ints) Any(f func(int64) bool) bool {
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

// =============== average.go =================

// Average returns the average of all elements in the slice.
// Can be generated for all number types.
func (s Ints) Average() float64 {
	var sum int64
	if len(s) == 0 {
		return float64(0)
	}
	for _, i := range s {
		sum += i
	}
	return float64(sum) / float64(len(s))
}

// =============== delete.go =================

// Delete returns a slice with the first matching element
// removed from the slice.
// Can be generated for any type.
func (s Ints) Delete(e int64) (out Ints) {
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

// =============== drop.go =================

func (s Ints) Drop(i int) (out Ints) {
	for i < len(s) {
		out = append(out, s[i])
		i++
	}
	return
}

// =============== elem.go =================

// Elem returns true if the slice contains the element
// Can be generated for any type.
func (s Ints) Elem(el int64) bool {
	for _, e := range s {
		if e == el {
			return true
		}
	}
	return false
}

// =============== filter.go =================

// Filter returns a slice containing only the elements that match the predicate.
// Can be generated for any type.
func (s Ints) Filter(f func(int64) bool) (out Ints) {
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}

// =============== foldl.go =================

// Foldl reduces a list by iteratively applying f from left->right. Thus, for an empty slice, the result is the default zero-value.
func (s Ints) Foldl(z int64, f func(e1, e2 int64) int64) (out int64) {
	if len(s) == 0 {
		return
	}
	out = s[0]
	for _, v := range s[1:] {
		out = f(out, v)
	}
	return f(out, z)
}

// =============== foldl1.go =================

// Foldl1 reduces a list by iteratively applying f from left->right. Thus, for an empty slice, the result is the default zero-value.
func (s Ints) Foldl1(f func(e1, e2 int64) int64) (out int64) {
	if len(s) == 0 {
		return
	}
	out = s[0]
	for _, v := range s[1:] {
		out = f(out, v)
	}
	return
}

// =============== head.go =================

// Head returns the first element in the slice.
// If no element is found, returns the zero-value of the type.
// Can be generated for any type.
func (s Ints) Head() (out int64) {
	if len(s) > 0 {
		out = s[0]
	}
	return
}

// =============== init.go =================

// Init takes n-1 elements from a slice, where n = len(list).
// Can be generated for any type.
func (s Ints) Init() (out Ints) {
	if len(s) == 0 {
		return
	}
	slicecopy := append([]int64(nil), s...)
	return slicecopy[:len(s)-1]
}

// =============== intercalate.go =================

// Intercalate inserts the method receiver slice into the function slice at each step.
// Can be generated for any type.
func (s Ints) Intercalate(ss [][]int64) (out Ints) {
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

// =============== intersperse.go =================

// Intersperse inserts the receiving value between each element of the method receiver.
// Can be generated for any type.
func (s Ints) Intersperse(value int64) (out Ints) {
	for i, el := range s {
		out = append(out, el)
		if i == len(s)-1 {
			break
		}
		out = append(out, value)
	}
	return
}

// =============== last.go =================

// Last returns the last element in the slice
// If no element is found, returns the zero-value of the type
// Can be generated for any type.
func (s Ints) Last() (out int64) {
	if len(s) > 0 {
		out = s[len(s)-1]
	}
	return
}

// =============== length.go =================

// Length returns the length (len) of a slice.
// Can be generated for any type.
func (s Ints) Length() int {
	return len(s)
}

// =============== map.go =================

// Map return a new slice with the map operation applied to each element.
// Can be generated for any type.
func (s Ints) Map(f func(int64) int64) (out Ints) {
	if f == nil {
		return s
	}
	for _, v := range s {
		out = append(out, f(v))
	}
	return
}

// =============== maximum.go =================

// Maximum returns the maximum in a slice.
// Can be generated for number types.
func (s Ints) Maximum() (out int64) {
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

// =============== maximumby.go =================

// MaximumBy returns the maximum elements according to a custom comparator.
// Can be generated for any type.
func (s Ints) MaximumBy(f func(e1, e2 int64) int64) (max int64) {
	if len(s) == 0 {
		return
	}
	max = s[0]
	for _, el := range s[1:] {
		max = f(max, el)
	}
	return
}

// =============== minimum.go =================

// Minimum returns the minimum of a slice.
// Can be generated for number types.
func (s Ints) Minimum() int64 {
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

// =============== modes.go =================

// Modes returns the elements with highest frequency in the slice.
// Can be generated for any type.
func (s Ints) Modes() (out Ints) {
	if len(s) == 0 {
		return
	}

	counts := make(map[int64]int)
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

// =============== nub.go =================

// Nub returns a slice containing only the unique elements of the receiver.
// The order of the elements is preserved.
// Can be generated for any type.
func (s Ints) Nub() (out Ints) {
	if len(s) == 0 {
		return
	}

	contains := make(map[int64]struct{})
	for _, v := range s {
		if _, ok := contains[v]; !ok {
			contains[v] = struct{}{}
			out = append(out, v)
		}
	}
	return
}

// =============== null.go =================

// Null returns true the slice is empty.
// Can be generated for any type.
func (s Ints) Null() bool {
	return len(s) == 0
}

// =============== product.go =================

// Product returns the product of all elements in the slice.
// Can be generated for any number type.
func (s Ints) Product() int64 {
	var prod int64
	for _, v := range s {
		prod += v
	}
	return prod
}

// =============== reverse.go =================

// Reverse returns the reversed slice.
// Can be generated for any type.
func (s Ints) Reverse() (out Ints) {
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return
}

// =============== sort.go =================

// Sort is a wrapper around go sort function.
// Can be generated for any type.
func (s Ints) Sort() Ints {
	out := make(Ints, len(s))
	copy(out, s)
	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})
	return out
}

// =============== span.go =================

// Span returns a tuple of any elements that satisfy the predicate up until the first failure, followed by
// the rest of the elements.
// Can be generated for any type.
func (s Ints) Span(f func(int64) bool) (before Ints, after Ints) {
	if f == nil {
		return before, s
	}

	failed := false

	for _, v := range s {
		if failed || !f(v) {
			after = append(after, v)
			failed = true
		} else {
			before = append(before, v)
		}
	}

	return
}

// =============== sum.go =================

// Sum returns the sum of all elements in the slice.
// Can be generated for any number type.
func (s Ints) Sum() int64 {
	var sum int64
	for _, v := range s {
		sum += v
	}
	return sum
}

// =============== tail.go =================

// Tail takes [1 -> n] elements from a slice, where n = len(list)
// Returns an empty slice if there are less than 2 elements in slice
// Can be generated for any type.
func (s Ints) Tail() (out Ints) {
	if len(s) <= 1 {
		return
	}
	slicecopy := append([]int64(nil), s...)
	return slicecopy[1:]
}

// =============== take.go =================

// Take takes the first n elements of the slice, or the entire slice if n > len(slice).
// Can be generated for any type.
func (s Ints) Take(n uint64) (out Ints) {
	if len(s) == 0 {
		return
	}
	out = make(Ints, len(s))
	copy(out, s)
	if n < uint64(len(s)) {
		return out[:n]
	}
	return
}

// =============== takewhile.go =================

// TakeWhile continues appending to the output as long as the predicate is satisfied.
// Can be generated for any type.
func (s Ints) TakeWhile(p func(int64) bool) (out Ints) {
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

// =============== uncons.go =================

// Uncons decomposes a slice into the head and tail component.
// Can be generated for any type.
func (s Ints) Uncons() (head int64, tail Ints) {
	return s.Head(), s.Tail()
}

// =============== unlines.go =================

// Unlines joins together the string representation of the slice with newlines after each element.
// Can be generated for any type.
func (s Ints) Unlines() (out string) {
	for i, v := range s {
		out += fmt.Sprintf("%v", v)
		if i != len(s)-1 {
			out += "\n"
		}
	}
	return
}

// =============== unwords.go =================

// Unwords joins together the string representation of the slice with newlines after each element.
// Can be generated for any type.
func (s Ints) Unwords() (out string) {
	for i, v := range s {
		out += fmt.Sprintf("%v", v)
		if i != len(s)-1 {
			out += " "
		}
	}
	return
}
