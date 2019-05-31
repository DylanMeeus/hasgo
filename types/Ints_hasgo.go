// code generated by hasgo. [DO NOT EDIT!]
package types

import (
	"math"
	"sort"
)

// =============== abs.go =================

func (s Ints) Abs() (out Ints) {
	for _, v := range s {
		out = append(out, int64(math.Abs(float64(v))))
	}
	return
}

// =============== all.go =================

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

// =============== filter.go =================

func (s Ints) Filter(f func(int64) bool) (out Ints) {
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}

// =============== head.go =================

// Returns the first element in the slice
// If no element is found, returns the zero-value of the type
func (s Ints) Head() (out int64) {
	if len(s) > 0 {
		out = s[0]
	}
	return
}

// =============== init.go =================

// Take n-1 elements from a slice, where n = len(list)
func (s Ints) Init() (out Ints) {
	slicecopy := append([]int64(nil), s...)
	return slicecopy[:len(s)-1]
}

// =============== intercalate.go =================

// inserts the method receiver slice into the function slice at each step
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

// =============== last.go =================

// Returns the last element in the slice
// If no element is found, returns the zero-value of the type
func (s Ints) Last() (out int64) {
	if len(s) > 0 {
		out = s[len(s)-1]
	}
	return
}

// =============== length.go =================

func (s Ints) Length() int {
	return len(s)
}

// =============== map.go =================

// Return a new slice with the map operation applied to each element.
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

// =============== minimum.go =================

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

// =============== null.go =================

// tests if the slice is empty
func (s Ints) Null() bool {
	return len(s) == 0
}

// =============== product.go =================

func (s Ints) Product() int64 {
	var prod int64
	for _, v := range s {
		prod += v
	}
	return prod
}

// =============== reverse.go =================

// Returns the reversed slice
func (s Ints) Reverse() (out Ints) {
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return
}

// =============== sort.go =================

// wrapper around go sort functions
func (s Ints) Sort() Ints {
	out := make(Ints, len(s))
	copy(out, s)
	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})
	return out
}

// =============== sum.go =================

func (s Ints) Sum() int64 {
	var sum int64
	for _, v := range s {
		sum += v
	}
	return sum
}

// =============== tail.go =================

// Take [1 -> n] elements from a slice, where n = len(list)
// Returns an empty slice if there are less than 2 elements in slice
func (s Ints) Tail() (out Ints) {
	if len(s) <= 1 {
		return
	}
	slicecopy := append([]int64(nil), s...)
	return slicecopy[1:]
}

// =============== take.go =================

func (s Ints) Take(n uint64) (out Ints) {
	out = make(Ints, len(s))
	copy(out, s)
	return out[:n]
}

// =============== uncons.go =================

func (s Ints) Uncons() (head int64, tail Ints) {
	return s.Head(), s.Tail()
}
