package functions

func (s SliceType) Take(n uint64) (out SliceType) {
	out = make(SliceType, len(s))
	copy(out, s)
	return out[:n]
}
