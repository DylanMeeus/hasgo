package types

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

// EqualsOrdered verifies that both slices contain the same elements in the same position.
func (is Strings) EqualsOrdered(other Strings) bool {
	if len(is) != len(other) {
		return false
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
