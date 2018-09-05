package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// EARLY SYSTEM:
	// t := time.Now()
	// fmt.Println(rand.Intn(t.Nanosecond()))
	// fmt.Println(rand.Intn(t.Nanosecond()))
	// fmt.Println(rand.Intn(t.Nanosecond()))

	// Set seed based by time
	rand.Seed(time.Now().UTC().UnixNano())

	// While loop test
	test := 0
	for test <= 10 {
		fmt.Print(rng(1, 10), " ")
		test++
	}
}

// RNG function
func rng(min int, max int) int {
	return min + rand.Intn(max-min)
}
