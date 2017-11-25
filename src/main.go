package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
)

func calculatePresents(desk uint32) uint32 {
	var total, i uint32
	total = 10 * desk
	for i = desk-1; i > 0; i-- {
		if desk % i == 0 {
			total += i * 10
		}
	}

	return total
}


func main() {
	var desk, target, currentMax, presentsForDesk, i uint32
	var start time.Time
	var duration time.Duration

	parsedTarget, _ := strconv.ParseInt(os.Getenv("PRESENTS"), 10, 64)
	target = uint32(parsedTarget)
	currentMax = 0

	start = time.Now()

	for i = 0; currentMax < target; i+=12 {
		presentsForDesk = calculatePresents(i)
		if presentsForDesk > currentMax {
			fmt.Printf("%v: %v\n", i, presentsForDesk)
			currentMax = presentsForDesk
			desk = i
		}
	}

	duration = time.Since(start)

	fmt.Printf("%v\n", desk)
	fmt.Printf("%v\n", duration)
}
