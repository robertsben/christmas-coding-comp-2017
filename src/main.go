package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
)

func main() {
	//var desk, primeTarget, currentMax, presentsForDesk, step uint32
	var desk, j, searchLimit, limit uint32
	var divisor uint8
	var start time.Time
	var duration time.Duration

	parsedTarget, _ := strconv.ParseInt(os.Getenv("PRESENTS"), 10, 64)

	// start timing now we have all the information
	start = time.Now()
	limit = uint32(parsedTarget)
	
	divisor = 10

	if limit >= 390 && limit < 20160 {
		divisor = 20
	} else if limit >= 20160 && limit <= 6770400 {
		divisor = 30
	} else if limit > 6770400 {
		divisor = 40
	}

	searchLimit = limit/uint32(divisor)

	cache := make([]uint32, searchLimit+1, searchLimit+1)

	for desk = 1; desk <= searchLimit; desk++ {
		for j = desk; j <= searchLimit; j += desk {
			cache[j] += desk
		}
		if cache[desk] * 10 > limit {
			break
		}
		//if cache[i] > currentMax {
		//	currentMax = cache[i]
		//	fmt.Printf("%v: %v, %v\n", i, cache[i], math.Floor(float64(cache[i])/float64(i)))
		//}
	}

	// quit timing now we have the result
	duration = time.Since(start)

	fmt.Printf("%v\n", desk)
	fmt.Printf("%v\n", duration)
}
