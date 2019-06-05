package types

import (
	"testing"
)

var (
	dylan = person{"Dylan", "Meeus", 26}
	ana   = person{"Ana", "Esparza", 30}
	sean  = person{"Sean", "West", 25}
	tom   = person{"Tom", "Devrolijke", 26}
	chris = person{"Chris", "Evans", 27}

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

	structPredicateTests = []struct {
		input     persons
		predicate func(person) bool
		any       bool
		all       bool
	}{
		{
			persons{dylan, chris},
			func(p person) bool { return p.age == 26 },
			true,
			false,
		},
		{
			persons{dylan, tom},
			func(p person) bool { return p.age == 26 },
			true,
			true,
		},
		{
			persons{dylan, tom},
			func(p person) bool { return p.age == 30 },
			false,
			false,
		},
	}

	structTakeTests = []struct {
		input      persons
		takeAmount uint64
		take       persons
		init       persons
		tail       persons
		head       person
		last       person
	}{
		{
			persons{dylan, ana, sean, tom},
			2,
			persons{dylan, ana},
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

	structLengthTests = []struct {
		input  persons
		output int
	}{
		{
			persons{dylan, ana, sean},
			3,
		},
		{
			persons{},
			0,
		},
		{
			nil,
			0,
		},
	}

	structBoolTests = []struct {
		input persons
		null  bool
	}{
		{
			persons{dylan},
			false,
		},
		{
			persons{},
			true,
		},
		{
			nil,
			true,
		},
	}

	structIntercalateTests = []struct {
		input       persons
		intercalate [][]person
		output      persons
	}{
		{
			persons{dylan},
			[][]person{
				{ana, sean},
				{tom},
			},
			persons{ana, sean, dylan, tom},
		},
		{
			persons{chris},
			[][]person{
				{ana, dylan},
				{sean},
				{tom},
			},
			persons{ana, dylan, chris, sean, chris, tom},
		},
	}

	structModesTests = []struct {
		input  persons
		output persons
	}{
		{
			persons{},
			persons{},
		},
		{
			persons{ana, dylan},
			persons{ana, dylan},
		},
		{
			persons{ana, ana, dylan, dylan},
			persons{ana, dylan},
		},
		{
			persons{ana, dylan, dylan},
			persons{dylan},
		},
	}

	structNubTests = []struct {
		input  persons
		output persons
	}{
		{
			nil,
			persons{},
		},
		{
			persons{},
			persons{},
		},
		{
			persons{ana, dylan},
			persons{ana, dylan},
		},
		{
			persons{ana, dylan, chris, dylan, ana},
			persons{ana, dylan, chris},
		},
		{
			persons{ana, ana, ana},
			persons{ana},
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

func Test_StructAny(t *testing.T) {
	for _, test := range structPredicateTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Any(test.predicate); res != test.any {
				t.Errorf("expected %v but got %v", test.any, res)
			}
		})
	}
}

func Test_StructAll(t *testing.T) {
	for _, test := range structPredicateTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.All(test.predicate); res != test.all {
				t.Errorf("expected %v but got %v", test.all, res)
			}
		})
	}
}

func Test_StructTake(t *testing.T) {
	for _, test := range structTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Take(test.takeAmount); !res.EqualsOrdered(test.take) {
				t.Errorf("expected %v but got %v", test.take, res)
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

func Test_StructLength(t *testing.T) {
	for _, test := range structLengthTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Length(); res != test.output {
				t.Errorf("expected %v but got %v", test.output, res)
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

func Test_StructNull(t *testing.T) {
	for _, test := range structBoolTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Null(); res != test.null {
				t.Errorf("expected %v but got %v", test.null, res)
			}
		})
	}
}

func Test_StructIntercalate(t *testing.T) {
	for _, test := range structIntercalateTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Intercalate(test.intercalate); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StructModes(t *testing.T) {
	for _, test := range structModesTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Modes(); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StructNub(t *testing.T) {
	for _, test := range intsNubTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Nub(); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}
