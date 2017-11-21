package main

import (
	"fmt"
	"time"
)

func calculatePresents(desk int) int {
	total := 10 * desk
	for i := desk-1; i > 0; i-- {
		if desk % i == 0 {
			total += i * 10
		}
	}

	return total
}


func main() {
	start := time.Now()
	fmt.Printf("Hello, world\n")

	target := 50000000
	currentMax := 0
	var desk int

	for i := 1; currentMax < target; i++ {
		presentsForDesk := calculatePresents(i)
		fmt.Printf("%v: %v\n", i, presentsForDesk)
		if presentsForDesk > currentMax {
			currentMax = presentsForDesk
			desk = i
		}
	}


	execTime := time.Since(start)
	fmt.Printf("%v\n", desk)
	fmt.Printf("%v\n", execTime)
}