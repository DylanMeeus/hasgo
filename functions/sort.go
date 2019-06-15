package functions

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
