package main

import (
	"fmt"
	hasgo "github.com/DylanMeeus/hasgo/types"
)

func main() {
	numbers := hasgo.Ints{}
	var i int64
	for i = 0; i < 1000; i++ {
		numbers = append(nums, i)
	}
	result := numbers.Filter(func(int64)bool{
		return i % 3 == 0 || i % 5 == 0
	}).Sum()
	fmt.Printf("%v\n", result)

}
