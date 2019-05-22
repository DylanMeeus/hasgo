package types

import (
	"testing"
)

var (
	dylan = person{"Dylan", "Meeus", 26}
	ana = person{"Ana", "Esparza", 30}
	sean = person{"Sean", "West", 25}
	tom = person{"Tom", "Devrolijke", 26}

	structFilterTests = []struct{
		input persons
		filter func(person)bool
		output persons
	}{
		{
			persons{dylan,ana,sean,tom},
			func(p person) bool { return p.age == 26 },
			persons{tom, dylan},
		},
	}

	structTakeTests = []struct{
		input persons
		init persons
		tail persons
		head person
		last person
	}{
		{
			persons{dylan,ana,sean,tom},
			persons{dylan,ana,sean},
			persons{ana,sean,tom},
			dylan,
			tom,
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
			if res := test.input.Init(); !res.Equals(test.init) {
				t.Errorf("expected %v but got %v", test.init, res)
			}
		})
	}
}


func Test_StructTail(t *testing.T) {
	for _, test := range structTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Tail(); !res.Equals(test.tail) {
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
