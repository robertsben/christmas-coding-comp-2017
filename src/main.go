package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
)

func main() {
	var desk, elf, searchLimit, limit uint32
	var start time.Time
	var duration time.Duration

	parsedTarget, _ := strconv.ParseInt(os.Getenv("PRESENTS"), 10, 64)
	limit = uint32(parsedTarget)

	// start timing now we have all the information
	start = time.Now()

	// figure out how small we can make our search range, based on patterns observed from spitting out results
	if limit >= 390 && limit < 20160 {
		searchLimit = limit/20
	} else if limit >= 20160 && limit <= 6770400 {
		searchLimit = limit/30
	} else if limit > 6770400 {
		searchLimit = limit/40
	} else {
		searchLimit = limit/10
	}

	// make a cache of desks, in which we'll store the quantity of presents
	cache := make([]uint32, searchLimit+1, searchLimit+1)

	// iterate desks up to the search limit
	for desk = 1; desk <= searchLimit; desk++ {
		// let each elf do his/her desk drop offs
		for elf = desk; elf <= searchLimit; elf += desk {
			cache[elf] += desk
		}
		// bail as soon as we find the desk with more than PRESENTS
		if cache[desk] * 10 >= limit {
			break
		}
	}

	// quit timing now we have the result
	duration = time.Since(start)

	fmt.Printf("%v\n", desk)
	fmt.Printf("%v\n", duration)
}
