package main

import (
	"fmt"
	hasgo "github.com/DylanMeeus/hasgo/types"
)

func main() {
	numbers := hasgo.Ints([]int64{1, 2, 3, 4})
	fmt.Printf("%v\n", numbers.Sum())
}
