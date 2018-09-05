package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	test := 0
	for test <= 10 {
		fmt.Print(rng(1, 10), " ")
		test++
	}
}

func rng(min int, max int) int {
	return min + rand.Intn(max-min)
}
