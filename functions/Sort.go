package functions

import (
	"sort"
)

// wrapper around go sort functions
func (s SliceType) Sort() SliceType {
	out := make(SliceType, len(s))
	copy(out, s)
	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})
	return out
}
