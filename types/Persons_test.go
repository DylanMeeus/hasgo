package types

import (
	"testing"
)

var (
	dylan = person{"Dylan", "Meeus", 26}
	ana   = person{"Ana", "Esparza", 30}
	sean  = person{"Sean", "West", 25}
	tom   = person{"Tom", "Devrolijke", 26}

	structFilterTests = []struct {
		input  persons
		filter func(person) bool
		output persons
	}{
		{
			persons{dylan, ana, sean, tom},
			func(p person) bool { return p.age == 26 },
			persons{tom, dylan},
		},
	}

	structTakeTests = []struct {
		input persons
		init  persons
		tail  persons
		head  person
		last  person
	}{
		{
			persons{dylan, ana, sean, tom},
			persons{dylan, ana, sean},
			persons{ana, sean, tom},
			dylan,
			tom,
		},
	}

	structReverseTests = []struct {
		input  persons
		output persons
	}{
		{
			persons{dylan, ana, sean, tom},
			persons{tom, sean, ana, dylan},
		},
	}

	structUnconsTests = []struct {
		input persons
		head  person
		tail  persons
	}{
		{
			persons{dylan, ana, sean, tom},
			dylan,
			persons{ana, sean, tom},
		},
	}

	structMapTests = []struct {
		input   persons
		mapfunc func(person) person
		output  persons
	}{
		{
			persons{dylan, ana, sean},
			func(p person) person { return dylan },
			persons{dylan, dylan, dylan},
		},
	}
)

func Test_StructFilter(t *testing.T) {
	for _, test := range structFilterTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Filter(test.filter); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StructInit(t *testing.T) {
	for _, test := range structTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Init(); !res.EqualsOrdered(test.init) {
				t.Errorf("expected %v but got %v", test.init, res)
			}
		})
	}
}

func Test_StructTail(t *testing.T) {
	for _, test := range structTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Tail(); !res.EqualsOrdered(test.tail) {
				t.Errorf("expected %v but got %v", test.tail, res)
			}
		})
	}
}

func Test_StructHead(t *testing.T) {
	for _, test := range structTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Head(); res != test.head {
				t.Errorf("expected %v but got %v", test.head, res)
			}
		})
	}
}

func Test_StructLast(t *testing.T) {
	for _, test := range structTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Last(); res != test.last {
				t.Errorf("expected %v but got %v", test.last, res)
			}
		})
	}
}

func Test_StructReverse(t *testing.T) {
	for _, test := range structReverseTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Reverse(); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StructUncons(t *testing.T) {
	for _, test := range structUnconsTests {
		t.Run("", func(t *testing.T) {
			if head, tail := test.input.Uncons(); head != test.head || !tail.EqualsOrdered(test.tail) {
				t.Errorf("expected (%v,%v) but got (%v,%v)", test.head, test.tail, head, tail)
			}
		})
	}
}

func Test_StructMap(t *testing.T) {
	for _, test := range structMapTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Map(test.mapfunc); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}
