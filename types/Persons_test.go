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
		input       persons
		takeAmount  uint64
		take        persons
		init        persons
		tail        persons
		head        person
		last        person
		whilePred   func(person) bool
		whileOutput persons
	}{
		{
			persons{dylan, ana, sean, tom},
			2,
			persons{dylan, ana},
			persons{dylan, ana, sean},
			persons{ana, sean, tom},
			dylan,
			tom,
			func(p person) bool { return p.age < 30 },
			persons{dylan},
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

	structIntersperseTests = []struct {
		input       persons
		intersperse person
		output      persons
	}{
		{
			nil,
			dylan,
			persons{},
		},
		{
			persons{ana, dylan, sean},
			chris,
			persons{ana, chris, dylan, chris, sean},
		},
		{
			persons{ana},
			dylan,
			persons{ana},
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

	structStringTests = []struct {
		input   persons
		unlines string
		unwords string
	}{
		{
			nil,
			"",
			"",
		},
		{
			persons{},
			"",
			"",
		},
		{
			persons{dylan, ana},
			"Dylan Meeus\nAna Esparza",
			"Dylan Meeus Ana Esparza",
		},
	}

	structDeleteTests = []struct {
		input   persons
		element person
		output  persons
	}{
		{
			nil,
			ana,
			persons{},
		},
		{
			persons{},
			dylan,
			persons{},
		},
		{
			persons{dylan, ana, ana, dylan},
			ana,
			persons{dylan, ana, dylan},
		},
	}

	structMinMaxByTests = []struct {
		input      persons
		comparator func(p1, p2 person) person
		max        person
	}{
		{
			nil,
			func(p1, p2 person) person { return p1 },
			person{},
		},
		{
			persons{dylan, ana, sean},
			func(p1, p2 person) person {
				if p1.age > p2.age {
					return p1
				}
				return p2
			},
			ana,
		},
	}

	structFoldTests = []struct {
		input     persons
		init      person
		foldlfunc func(p1, p2 person) person
		foldl     person
		foldl1    person
	}{
		{
			nil,
			person{},
			func(p1, p2 person) person { return p1 },
			person{},
			person{},
		},
		{
			persons{dylan, ana, sean},
			person{},
			func(p1, p2 person) person {
				if p1.age > p2.age {
					return p1
				}
				return p2
			},
			ana,
			ana,
		},
	}

	structElemTests = []struct {
		input  persons
		needle person
		output bool
	}{
		{
			nil,
			ana,
			false,
		},
		{
			persons{},
			ana,
			false,
		},
		{
			persons{dylan, ana, sean},
			ana,
			true,
		},
		{
			persons{dylan, sean},
			ana,
			false,
		},
	}

	structDropTests = []struct {
		input  persons
		drop   int
		output persons
	}{
		{
			nil,
			0,
			persons{},
		},
		{
			persons{dylan, ana, sean, chris},
			5,
			persons{},
		},
		{
			persons{dylan, ana, sean, chris},
			1,
			persons{ana, sean, chris},
		},
		{
			persons{dylan, ana, sean, chris},
			3,
			persons{chris},
		},
	}

	structDropWhileTests = []struct {
		input    persons
		dropfunc func(person) bool
		output   persons
	}{
		{
			nil,
			func(p person) bool { return p.age < 27 },
			persons{},
		},
		{
			persons{dylan, ana, sean, tom, chris},
			nil,
			persons{dylan, ana, sean, tom, chris},
		},
		{
			persons{dylan, ana, sean, tom, chris},
			func(p person) bool { return p.age < 50 },
			persons{},
		},
		{
			persons{dylan, ana, sean, tom, chris},
			func(p person) bool { return p.age < 30 },
			persons{ana, sean, tom, chris},
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

func Test_StructTakeWhile(t *testing.T) {
	for _, test := range structTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.TakeWhile(test.whilePred); !res.EqualsOrdered(test.whileOutput) {
				t.Errorf("expected %v but got %v", test.whileOutput, res)
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

func Test_StructIntersperse(t *testing.T) {
	for _, test := range structIntersperseTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Intersperse(test.intersperse); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
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
	for _, test := range structNubTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Nub(); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StructUnlines(t *testing.T) {
	for _, test := range structStringTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Unlines(); res != test.unlines {
				t.Errorf("expected %v but got %v", test.unlines, res)
			}
		})
	}
}

func Test_StructUnwords(t *testing.T) {
	for _, test := range structStringTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Unwords(); res != test.unwords {
				t.Errorf("expected %v but got %v", test.unwords, res)
			}
		})
	}
}

func Test_StructDelete(t *testing.T) {
	for _, test := range structDeleteTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Delete(test.element); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_structMaximumBy(t *testing.T) {
	for _, test := range structMinMaxByTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.MaximumBy(test.comparator); res != test.max {
				t.Errorf("expected %v but got %v", test.max, res)
			}
		})
	}
}

func Test_structFoldl(t *testing.T) {
	for _, test := range structFoldTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Foldl(test.init, test.foldlfunc); res != test.foldl {
				t.Errorf("expected %v but got %v", test.foldl, res)
			}
		})
	}
}

func Test_structFoldl1(t *testing.T) {
	for _, test := range structFoldTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Foldl1(test.foldlfunc); res != test.foldl1 {
				t.Errorf("expected %v but got %v", test.foldl1, res)
			}
		})
	}
}

func Test_structElem(t *testing.T) {
	for _, test := range structElemTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Elem(test.needle); res != test.output {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_structDrop(t *testing.T) {
	for _, test := range structDropTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Drop(test.drop); !test.output.Equals(res) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_structDropWhile(t *testing.T) {
	for _, test := range structDropWhileTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.DropWhile(test.dropfunc); !test.output.Equals(res) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}
