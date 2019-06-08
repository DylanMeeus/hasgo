package types

//go:generate hasgo -T=person -S=persons
type persons []person

type person struct {
	firstname string
	lastname  string
	age       int
}

func (p person) String() string {
	return p.firstname + " " + p.lastname
}

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
