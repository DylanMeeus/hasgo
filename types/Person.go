package types

//go:generate hasgo -T=person -S=persons
type persons []person

type person struct {
	firstname string
	lastname  string
	age       int
}

// String returns the firstname and lastname of the person.
func (p person) String() string {
	return p.firstname + " " + p.lastname
}

// EqualsOrdered verifies that both slices have the same elements in the same position.
func (is persons) EqualsOrdered(other persons) bool {
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

// Equals verifies that both slices have the same elements regardless of position.
func (is persons) Equals(other persons) bool {
	if len(is) != len(other) {
		return false
	}
	contains := make(map[person]struct{}, 0)
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
