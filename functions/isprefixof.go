package functions

// IsPrefixOf returns true if the current sliceType is a prefix of the passed one.
// Can be generated for any time.
func (s SliceType) IsPrefixOf(in SliceType) bool {
	if len(s) > len(in) {
		return false
	}
	for i, v := range s {
		if in[i] != v {
			return false
		}
	}
	return true
}
