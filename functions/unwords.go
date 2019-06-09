package functions

import (
	"fmt"
)

// Joins together the string representation of the slice
// With newlines after each element.
func (s SliceType) Unwords() (out string) {
	for i, v := range s {
		out += fmt.Sprintf("%v", v)
		if i != len(s)-1 {
			out += " "
		}
	}
	return
}
