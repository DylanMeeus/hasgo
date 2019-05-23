package main

import (
	"fmt"
	. "github.com/DylanMeeus/hasgo/types"
)

// these are just examples, not the way you'd actually implement them in Go.

func main() {
	fmt.Printf("product: %v\n", Multiply(1, Ints{1,2,3}))
	EpicFunction()
}


// from the docs
func EpicFunction() {
	result := IntRange(-10, 10).
		Abs().
		Filter(func(i int64) bool {
			return i % 2 == 0
		}).
		Sum()
	fmt.Printf("%v\n", result)
}

// Recursively multiply numbers
func Multiply(product int64, numbers Ints) int64 {
	if len(numbers) == 0 {
		return product 
	}
	h, rest := numbers.Uncons()
	return Multiply(product * h, rest)
}
