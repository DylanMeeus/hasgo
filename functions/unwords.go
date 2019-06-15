package functions

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
