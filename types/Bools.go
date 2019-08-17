package types

// Bools represents an array of booleans
type Bools []bool

// And returns true if all bools are true
func (bs Bools) And() bool {
	if len(bs) == 0 {
		return true
	}
	for _, b := range bs {
		if !b {
			return false
		}
	}
	return true
}

// Or returns true if any value is true
func (bs Bools) Or() bool {
	if len(bs) == 0 {
		return true
	}
	for _, b := range bs {
		if b {
			return true
		}
	}
	return false
}
