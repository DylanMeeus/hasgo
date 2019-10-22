package types

import (
	"testing"
)

// unit testing the Ints
var (
	intsMathTests = []struct {
		input   Ints
		sum     int64
		product int64
		average float64
	}{
		{
			nil,
			0,
			0,
			0,
		},
		{
			Ints([]int64{}),
			0,
			0,
			0,
		},
		{
			Ints([]int64{1, 2, 3}),
			6,
			6,
			2,
		},
	}

	intsFilterTests = []struct {
		input  Ints
		filter func(int64) bool
		output Ints
	}{
		{
			nil,
			nil,
			Ints{},
		},
		{
			Ints{},
			nil,
			Ints{},
		},
		{
			Ints{1, 2, 3, 4},
			func(i int64) bool { return i%2 == 0 },
			Ints{2, 4},
		},
	}

	intsPredicateTests = []struct {
		input     Ints
		predicate func(int64) bool
		any       bool
		all       bool
	}{
		{
			nil,
			nil,
			false,
			false,
		},
		{
			Ints{},
			nil,
			false,
			false,
		},
		{
			Ints{1, 2, 3, 4},
			func(i int64) bool { return i%2 == 0 },
			true,
			false,
		},
		{
			Ints{2, 4, 6, 8},
			func(i int64) bool { return i%2 == 0 },
			true,
			true,
		},
		{
			Ints{1, 3, 5, 7},
			func(i int64) bool { return i%2 == 0 },
			false,
			false,
		},
	}

	intsRangeTests = []struct {
		start  int64
		stop   int64
		output Ints
	}{
		{0, 10, Ints{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{-10, 1, Ints{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1}},
	}

	intsReplicateTests = []struct {
		count  uint64
		value  int64
		output Ints
	}{
		{5, 1337, Ints{1337, 1337, 1337, 1337, 1337}},
	}

	// tests that take a subset of the slice
	intsTakeTests = []struct {
		input       Ints
		takeAmount  uint64
		take        Ints
		init        Ints
		tail        Ints
		head        int64
		last        int64
		whilePred   func(int64) bool
		whileOutput Ints
	}{
		{
			nil,
			3,
			Ints{},
			Ints{},
			Ints{},
			0,
			0,
			nil,
			Ints{},
		},
		{
			Ints{},
			3,
			Ints{},
			Ints{},
			Ints{},
			0,
			0,
			nil,
			Ints{},
		},
		{
			IntRange(0, 10),
			3,
			Ints{0, 1, 2},
			IntRange(0, 9),
			IntRange(1, 10),
			0,
			10,
			func(u int64) bool { return u < 5 },
			IntRange(0, 4),
		},
		{
			IntRange(0, 10),
			11,
			IntRange(0, 10),
			IntRange(0, 9),
			IntRange(1, 10),
			0,
			10,
			func(u int64) bool { return u < 100 },
			IntRange(0, 10),
		},
	}

	intsReverseTests = []struct {
		input  Ints
		output Ints
	}{
		{
			Ints{},
			Ints{},
		},
		{
			nil,
			Ints{},
		},
		{
			Ints{1, 2, 3},
			Ints{3, 2, 1},
		},
	}

	intsUnconsTests = []struct {
		input Ints
		head  int64
		tail  Ints
	}{
		{
			Ints{1, 2, 3, 4},
			1,
			Ints{2, 3, 4},
		},
		{
			Ints{},
			0,
			Ints{},
		},
		{
			nil,
			0,
			Ints{},
		},
	}

	intsMinMaxTests = []struct {
		input Ints
		min   int64
		max   int64
	}{
		{
			nil,
			0,
			0,
		},
		{
			Ints{},
			0,
			0,
		},
		{
			Ints{-1, 5, 3, -100},
			-100,
			5,
		},
	}

	intsMinMaxByTests = []struct {
		input      Ints
		comparator func(i1, i2 int64) int64
		max        int64
	}{
		{
			nil,
			func(i1, i2 int64) int64 { return i1 },
			0,
		},
		{
			Ints{1, 2, 3, 4},
			func(i1, i2 int64) int64 {
				if i1 > i2 {
					return i1
				}
				return i2
			},
			4,
		},
	}

	intsMapTests = []struct {
		input   Ints
		mapfunc func(int64) int64
		output  Ints
	}{
		{
			nil,
			func(i int64) int64 { return i * i },
			Ints{},
		},
		{
			Ints{},
			func(i int64) int64 { return i * i },
			Ints{},
		},
		{
			Ints{1, 2, 3},
			func(i int64) int64 { return i * i },
			Ints{1, 4, 9},
		},
		{
			Ints{1, 2, 3},
			nil,
			Ints{1, 2, 3},
		},
	}

	intsSortTests = []struct {
		input  Ints
		output Ints
	}{
		{
			nil,
			Ints{},
		},
		{
			Ints{},
			Ints{},
		},
		{
			Ints{3, 2, 2, 1},
			Ints{1, 2, 2, 3},
		},
	}

	intsLengthTests = []struct {
		input  Ints
		output int
	}{
		{
			Ints{1, 2, 3},
			3,
		},
		{
			nil,
			0,
		},
		{
			Ints{},
			0,
		},
	}

	// tests that result in boolean
	intsBoolTests = []struct {
		input Ints
		null  bool
	}{
		{
			nil,
			true,
		},
		{
			Ints{},
			true,
		},
		{
			Ints{1, 2, 3},
			false,
		},
	}

	intsIntersperseTests = []struct {
		input       Ints
		intersperse int64
		output      Ints
	}{
		{
			nil,
			0,
			Ints{},
		},
		{
			nil,
			1337,
			Ints{},
		},
		{
			Ints{},
			0,
			Ints{},
		},
		{
			Ints{1, 2, 3},
			0,
			Ints{1, 0, 2, 0, 3},
		},
		{
			Ints{1, 2, 3},
			1337,
			Ints{1, 1337, 2, 1337, 3},
		},
	}

	intsIntercalateTests = []struct {
		input       Ints
		intercalate [][]int64
		output      Ints
	}{
		{
			nil,
			[][]int64{
				{4, 4},
				{5, 5},
			},
			Ints{4, 4, 5, 5},
		},
		{
			Ints{},
			[][]int64{
				{4, 4},
				{5, 5},
			},
			Ints{4, 4, 5, 5},
		},
		{
			Ints{1, 2},
			[][]int64{
				{4, 4},
				{5, 5},
			},
			Ints{4, 4, 1, 2, 5, 5},
		},
		{
			Ints{1, 2},
			[][]int64{
				{4, 4},
				{5, 5},
				{6, 6},
			},
			Ints{4, 4, 1, 2, 5, 5, 1, 2, 6, 6},
		},
	}

	intsModesTests = []struct {
		input  Ints
		output Ints
	}{
		{
			nil,
			Ints{},
		},
		{
			Ints{},
			Ints{},
		},
		{
			Ints{1, 2, 3},
			Ints{1, 2, 3},
		},
		{
			Ints{1, 1, 2, 2, 3, 3},
			Ints{1, 2, 3},
		},
		{
			Ints{1, 2, 2, 3},
			Ints{2},
		},
		{
			Ints{1, 2, 2, 3, 3},
			Ints{2, 3},
		},
	}

	intsNubTests = []struct {
		input  Ints
		output Ints
	}{
		{
			nil,
			Ints{},
		},
		{
			Ints{},
			Ints{},
		},
		{
			Ints{1, 2, 3},
			Ints{1, 2, 3},
		},
		{
			Ints{1, 2, 3, 2, 1},
			Ints{1, 2, 3},
		},
		{
			Ints{3, 3, 3, 3, 2, 1, 2, 3},
			Ints{3, 2, 1},
		},
		{
			Ints{1, 1, 1, 1, 1},
			Ints{1},
		},
	}

	intsStringTests = []struct {
		input   Ints
		unlines string
		unwords string
	}{
		{
			nil,
			"",
			"",
		},
		{
			Ints{},
			"",
			"",
		},
		{
			Ints{1, 2, 3},
			"1\n2\n3",
			"1 2 3",
		},
	}

	intsDeleteTests = []struct {
		input   Ints
		element int64
		output  Ints
	}{
		{
			nil,
			0,
			Ints{},
		},
		{
			Ints{},
			0,
			Ints{},
		},
		{
			Ints{1, 2, 3, 3, 4, 5, 3},
			3,
			Ints{1, 2, 3, 4, 5, 3},
		},
		{
			Ints{1, 2, 3},
			-1,
			Ints{1, 2, 3},
		},
	}

	intsFoldTests = []struct {
		input    Ints
		init     int64
		foldfunc func(int64, int64) int64
		foldl    int64
		foldl1   int64
	}{
		{
			nil,
			0,
			func(i1, i2 int64) int64 { return i1 + i2 },
			0,
			0,
		},
		{
			Ints{1},
			1,
			func(i1, i2 int64) int64 { return i1 + i2 },
			2,
			1,
		},
		{
			Ints{1, 2, 3, 4, 5},
			0,
			func(i1, i2 int64) int64 { return i1 - i2 },
			-13,
			-13,
		},
	}

	intsElemTests = []struct {
		input  Ints
		needle int64
		output bool
	}{
		{
			nil,
			0,
			false,
		},
		{
			Ints{},
			0,
			false,
		},
		{
			Ints{1, 2, 3, 4},
			1,
			true,
		},
		{
			Ints{1, 2, 3, 4},
			100,
			false,
		},
	}

	intsDropTests = []struct {
		input  Ints
		drop   int
		output Ints
	}{
		{
			nil,
			0,
			Ints{},
		},
		{
			Ints{1, 2, 3, 4},
			5,
			Ints{},
		},
		{
			Ints{1, 2, 3, 4},
			1,
			Ints{2, 3, 4},
		},
		{
			Ints{1, 2, 3, 4},
			3,
			Ints{4},
		},
	}
)

func Test_IntsSum(t *testing.T) {
	for _, test := range intsMathTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Sum(); res != test.sum {
				t.Errorf("expected %v but got %v", test.sum, res)
			}
		})
	}
}

func Test_IntsProduct(t *testing.T) {
	for _, test := range intsMathTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Product(); res != test.product {
				t.Errorf("expected %v but got %v", test.product, res)
			}
		})
	}
}

func Test_IntsAverage(t *testing.T) {
	for _, test := range intsMathTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Average(); res != test.average {
				t.Errorf("expected %v but got %v", test.product, res)
			}
		})
	}
}

func Test_IntsFilter(t *testing.T) {
	for _, test := range intsFilterTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Filter(test.filter); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsAll(t *testing.T) {
	for _, test := range intsPredicateTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.All(test.predicate); res != test.all {
				t.Errorf("expected %v but got %v", test.all, res)
			}
		})
	}
}

func Test_IntsAny(t *testing.T) {
	for _, test := range intsPredicateTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Any(test.predicate); res != test.any {
				t.Errorf("expected %v but got %v", test.any, res)
			}
		})
	}
}

func Test_IntsRange(t *testing.T) {
	for _, test := range intsRangeTests {
		t.Run("", func(t *testing.T) {
			if res := IntRange(test.start, test.stop); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsReplicate(t *testing.T) {
	for _, test := range intsReplicateTests {
		t.Run("", func(t *testing.T) {
			if res := IntReplicate(test.count, test.value); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsTake(t *testing.T) {
	for _, test := range intsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Take(test.takeAmount); !res.EqualsOrdered(test.take) {
				t.Errorf("expected %v but got %v", test.take, res)
			}
		})
	}
}

func Test_IntsTakewhile(t *testing.T) {
	for _, test := range intsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.TakeWhile(test.whilePred); !res.EqualsOrdered(test.whileOutput) {
				t.Errorf("expected %v but got %v", test.whileOutput, res)
			}
		})
	}
}

func Test_IntsInit(t *testing.T) {
	for _, test := range intsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Init(); !res.EqualsOrdered(test.init) {
				t.Errorf("expected %v but got %v", test.init, res)
			}
		})
	}
}

func Test_IntsTail(t *testing.T) {
	for _, test := range intsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Tail(); !res.EqualsOrdered(test.tail) {
				t.Errorf("expected %v but got %v", test.tail, res)
			}
		})
	}
}

func Test_IntsHead(t *testing.T) {
	for _, test := range intsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Last(); res != test.last {
				t.Errorf("expected %v but got %v", test.last, res)
			}
		})
	}
}

func Test_IntsLast(t *testing.T) {
	for _, test := range intsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Head(); res != test.head {
				t.Errorf("expected %v but got %v", test.head, res)
			}
		})
	}
}

func Test_IntsReverse(t *testing.T) {
	for _, test := range intsReverseTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Reverse(); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsUncons(t *testing.T) {
	for _, test := range intsUnconsTests {
		t.Run("", func(t *testing.T) {
			if head, tail := test.input.Uncons(); head != test.head || !tail.EqualsOrdered(test.tail) {
				t.Errorf("expected (%v,%v) but got (%v,%v)", test.head, test.tail, head, tail)
			}
		})
	}
}

func Test_IntsMaximum(t *testing.T) {
	for _, test := range intsMinMaxTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Maximum(); res != test.max {
				t.Errorf("expected %v but got %v", test.max, res)
			}
		})
	}
}

func Test_IntsMinimum(t *testing.T) {
	for _, test := range intsMinMaxTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Minimum(); res != test.min {
				t.Errorf("expected %v but got %v", test.min, res)
			}
		})
	}
}

func Test_IntsMap(t *testing.T) {
	for _, test := range intsMapTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Map(test.mapfunc); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsSort(t *testing.T) {
	for _, test := range intsSortTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Sort(); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsLength(t *testing.T) {
	for _, test := range intsLengthTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Length(); res != test.output {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsNull(t *testing.T) {
	for _, test := range intsBoolTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Null(); res != test.null {
				t.Errorf("expected %v but got %v", test.null, res)
			}
		})
	}
}

func Test_IntsIntersperse(t *testing.T) {
	for _, test := range intsIntersperseTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Intersperse(test.intersperse); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsIntercalate(t *testing.T) {
	for _, test := range intsIntercalateTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Intercalate(test.intercalate); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsModes(t *testing.T) {
	for _, test := range intsModesTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Modes(); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsNub(t *testing.T) {
	for _, test := range intsNubTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Nub(); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsUnlines(t *testing.T) {
	for _, test := range intsStringTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Unlines(); res != test.unlines {
				t.Errorf("expected %v but got %v", test.unlines, res)
			}
		})
	}
}

func Test_IntsUnwords(t *testing.T) {
	for _, test := range intsStringTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Unwords(); res != test.unwords {
				t.Errorf("expected %v but got %v", test.unwords, res)
			}
		})
	}
}

func Test_IntsDelete(t *testing.T) {
	for _, test := range intsDeleteTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Delete(test.element); !test.output.EqualsOrdered(res) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsMaximumBy(t *testing.T) {
	for _, test := range intsMinMaxByTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.MaximumBy(test.comparator); res != test.max {
				t.Errorf("expected %v but got %v", test.max, res)
			}
		})
	}
}

func Test_IntsFoldl(t *testing.T) {
	for _, test := range intsFoldTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Foldl(test.init, test.foldfunc); res != test.foldl {
				t.Errorf("expected %v but got %v", test.foldl, res)
			}
		})
	}
}

func Test_IntsFoldl1(t *testing.T) {
	for _, test := range intsFoldTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Foldl1(test.foldfunc); res != test.foldl1 {
				t.Errorf("expected %v but got %v", test.foldl1, res)
			}
		})
	}
}

func Test_IntsElem(t *testing.T) {
	for _, test := range intsElemTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Elem(test.needle); res != test.output {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_IntsDrop(t *testing.T) {
	for _, test := range intsDropTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Drop(test.drop); !test.output.Equals(res) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}
