package functions

func (s SliceType) Take(n uint64) (out SliceType) {
	if len(s) == 0 {
		return
	}
	out = make(SliceType, len(s))
	copy(out, s)
	if n < uint64(len(s)) {
		return out[:n]
	}
	return
}
