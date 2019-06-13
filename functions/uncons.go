package functions

// Decompose a slice into the head and tail component.
// Can be generated for any type.
func (s SliceType) Uncons() (head ElementType, tail SliceType) {
	return s.Head(), s.Tail()
}
