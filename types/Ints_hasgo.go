// code generated by hasgo. [DO NOT EDIT!]
package types

import (
	"math"
)

//===============Tail.go=============

// Take [1 -> n] elements from a slice, where n = len(list)
// Returns an empty slice if there are less than 2 elements in slice
func (s Ints) Tail() (out Ints) {
	if len(s) <= 1 {
		return
	}
	slicecopy := append([]int64(nil), s...)
	return slicecopy[1:]
}

//===============Abs.go=============

func (s Ints) Abs() (out Ints) {
	for _, v := range s {
		out = append(out, int64(math.Abs(float64(v))))
	}
	return
}

//===============Filter.go=============

func (s Ints) Filter(f func(int64) bool) (out Ints) {
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}

//===============Head.go=============

// Returns the first element in the slice
// If no element is found, returns the zero-value of the type
func (s Ints) Head() (out int64) {
	if len(s) > 0 {
		out = s[0]
	}
	return
}

//===============Init.go=============

// Take n-1 elements from a slice, where n = len(list)
func (s Ints) Init() (out Ints) {
	slicecopy := append([]int64(nil), s...)
	return slicecopy[:len(s)-1]
}

//===============Last.go=============

// Returns the last element in the slice
// If no element is found, returns the zero-value of the type
func (s Ints) Last() (out int64) {
	if len(s) > 0 {
		out = s[len(s)-1]
	}
	return
}

//===============Sum.go=============

func (s Ints) Sum() int64 {
	var sum int64
	for _, v := range s {
		sum += v
	}
	return sum
}