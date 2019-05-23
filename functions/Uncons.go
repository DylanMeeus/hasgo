package functions

func (s SliceType) Uncons() (head ElementType, tail SliceType) {
	return s.Head(), s.Tail()
}
