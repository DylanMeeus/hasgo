package types

import (
	"strings"
)

// Strings is a wrapper around []string
//go:generate hasgo -T=string -S=Strings
type Strings []string

// StringReplicate creates a slice with `value` repeated `count` times
func StringReplicate(count uint64, value string) (out Strings) {
	out = make(Strings, count)
	for i := uint64(0); i < count; i++ {
		out[i] = value
	}
	return
}

const (
	wordSeparator = " "
	lineSeparator = "\n"
)

// EqualsOrdered verifies that both slices contain the same elements in the same position.
func (is Strings) EqualsOrdered(other Strings) bool {
	if len(is) != len(other) {
		return false
	}
	if len(is) == 0 && len(other) == 0 {
		return true
	}
	for i := range is {
		if is[i] != other[i] {
			return false
		}
	}
	return true
}

// Equals verifies that both slices contain the same elements, regardless of position.
func (is Strings) Equals(other Strings) bool {
	if len(is) != len(other) {
		return false
	}
	if len(is) == 0 && len(other) == 0 {
		return true
	}
	contains := make(map[string]struct{}, 0)
	for _, i := range is {
		contains[i] = struct{}{}
	}
	for _, i := range other {
		if _, ok := contains[i]; !ok {
			return false
		}
	}
	return true
}

// Words is a wrapper around strings.Split(string, " ")
func Words(s string) Strings {
	return split(s, wordSeparator)
}

// Lines is a wrapper around strings.Split(string, "\n")
func Lines(s string) Strings {
	return split(s, lineSeparator)
}

func split(s, sep string) Strings {
	if s == "" {
		return Strings{}
	}
	return Strings(strings.Split(s, sep))
}
