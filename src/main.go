package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
	//"math"
	"log"
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
	//var deskIsMoreThanSqrtOfPresents bool
	var start time.Time
	var duration time.Duration

	f, err := os.OpenFile("result-data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	parsedTarget, _ := strconv.ParseInt(os.Getenv("PRESENTS"), 10, 64)
	target = uint32(parsedTarget)
	currentMax = 0
	start = time.Now()

	fmt.Printf("%v!\n", target)
	for i = 1; currentMax < target; i+=1 {
		presentsForDesk = calculatePresents(i)
		//deskIsMoreThanSqrtOfPresents = i >= uint32(math.Sqrt(float64(presentsForDesk)))
		//fmt.Printf("%v: %v: %v\n", i, presentsForDesk, deskIsMoreThanSqrtOfPresents)
		if presentsForDesk > currentMax {
			currentMax = presentsForDesk
			desk = i
			fmt.Fprintf(f, "%v: %v\n", desk, presentsForDesk)
		}
	}

	duration = time.Since(start)

	fmt.Printf("%v\n", desk)
	fmt.Printf("%v\n", duration)
}
