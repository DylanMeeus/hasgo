package main

import (
	"fmt"
	hasgo "github.com/DylanMeeus/hasgo/types"
)

// these are just examples, not the way you'd actually implement them in Go.

func main() {
	fmt.Printf("product: %v\n", Multiply(1, hasgo.Ints{1, 2, 3}))
	EpicFunction()
	fmt.Printf("fold %v\n", fold())
}

// EpicFunction is a simple example as seen on the github homepage.
func EpicFunction() {
	result := hasgo.IntRange(-10, 10).
		Abs().
		Filter(func(i int64) bool {
			return i%2 == 0
		}).
		Sum()
	fmt.Printf("%v\n", result)
}

// Multiply implements a recursively multiplier
func Multiply(product int64, numbers hasgo.Ints) int64 {
	if len(numbers) == 0 {
		return product
	}
	h, rest := numbers.Uncons()
	return Multiply(product*h, rest)
}

func fold() int64 {
	result := hasgo.IntRange(0, 10).
		Foldr(0, func(i1, i2 int64) int64 { return i1 - i2 })
	return result

}
