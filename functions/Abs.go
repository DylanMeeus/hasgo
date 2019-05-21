package functions

import (
	"math"
)

func (s SliceType) Abs() (out SliceType) {
	for _, v := range s {
		out = append(out, ElementType(math.Abs(float64(v))))
	}
	return
}
