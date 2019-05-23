package main

import (
	"./player"
	"fmt"
)

func main() {
	p := player.Player{"Insanity", 1337}
	p2 := player.Player{"Unknown", 2041}
	ps := player.Players{p, p2}
	fmt.Printf("%v\n", ps.Head())
}
