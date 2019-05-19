package main

import (
	"fmt"
	hasgo "github.com/DylanMeeus/hasgo/types"
)

func main() {
	hasgoFlavour()
	vanillaGo()
}

// The hasgo way of solving this problem
func hasgoFlavour() {
	result := hasgo.IntRange(0,1000).Filter(func(i int64) bool {
		return i%3 == 0 || i%5 == 0
	}).Sum()
	fmt.Printf("%v\n", result)
}

// this is how you could solve this in normal Go code.
func vanillaGo() {
	var sum int
	for i := 0; i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	fmt.Printf("%v\n", sum)
}


