package functions

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
