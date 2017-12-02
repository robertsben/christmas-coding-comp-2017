package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
	"math"
)


type primeMultipleData struct {
	multiple, sum uint32 // p^x, sum of p^x + p^x-1 + ... p^0
}

type primeData struct {
	// here we assume we only need primes up to 65535; i.e. we'll deal only with numbers <= 4294836225 (almost all unsigned 32 bit integers)
	prime uint32 // p
	multiples []primeMultipleData
}

/* The minimum number of presents to find */
var limit uint32

/* A list of primes, with exponent and sum data */
var primes []primeData

/* Generate the exponents and sums of them */
func generateExponents(n uint32) []primeMultipleData {
	var i uint8
	var cumulativeExponent, nPower uint32
	var primeExponentList []primeMultipleData
	cumulativeExponent = 0
	primeExponentList = append(primeExponentList, primeMultipleData{1, 1})
	primeExponentList = append(primeExponentList, primeMultipleData{n, n + 1})

	for i = 2; uint64(cumulativeExponent) * uint64(n) <= uint64(limit); i++ {
		nPower = primeExponentList[i-1].multiple * n
		cumulativeExponent = primeExponentList[i-1].sum + nPower
		primeExponentList = append(primeExponentList, primeMultipleData{nPower, cumulativeExponent})
	}

	return primeExponentList
}

/* Generate the list of primes using atkins sieve */
func atkinSieve(max uint32) {
	var x, y, n uint32
	maxSqrt := math.Sqrt(float64(max))

	isPrime := make([]bool, max+1, max+1)

	for x = 1; float64(x) <= maxSqrt; x++ {
		for y = 1; float64(y) <= maxSqrt; y++ {
			n = (4*x*x) + (y*y)
			if n <= max && (n % 12 == 1 || n %12 == 5) {
				isPrime[n] = !isPrime[n]
			}
			n = (3*x*x) + (y*y)
			if n <= max && n % 12 == 7 {
				isPrime[n] = !isPrime[n]
			}
			n = (3*x*x) - (y*y)
			if x > y && n <= max && n % 12 == 11 {
				isPrime[n] = !isPrime[n]
			}
		}
	}

	for n = 5; float64(n) <= maxSqrt; n++ {
		if isPrime[n] {
			for y = n*n; y < max; y += n*n {
				isPrime[y] = false
			}
		}
	}

	isPrime[2] = true
	isPrime[3] = true

	for x = 0; int(x) < len(isPrime)-1; x++ {
		if isPrime[x] {
			multipleDataList := generateExponents(x)
			primes = append(primes, primeData{x, multipleDataList})
		}
	}
}

func calculatePresentsFromCache(desk uint32) uint32 {
	var total, primeTotal uint32
	var j uint16
	var prime primeData
	total = 1
	for _, prime = range primes {
		primeTotal = 1
		for j = 1; j < uint16(len(prime.multiples)) && desk % prime.multiples[j].multiple == 0; j++ {
			primeTotal = prime.multiples[j].sum
		}
		total *= primeTotal
	}
	// return total of exponent sums, times the elf present delivery multiplier
	return total * 10
}


func main() {
	var desk, primeTarget, currentMax, presentsForDesk, step uint32
	var start time.Time
	var duration time.Duration
	step = 1
	desk = 1

	parsedTarget, _ := strconv.ParseInt(os.Getenv("PRESENTS"), 10, 64)

	// start timing now we have all the information
	start = time.Now()
	limit = uint32(parsedTarget)
	primeTarget = uint32(math.Sqrt(float64(limit)))

	/* Desk is always > limit/50 (as long as the result >= 2) */
 	if limit >= 100 {
		desk = limit/50
		step = 2
		if desk > 2 && desk % 2 != 0 {
			desk--
		}
	}

	/* initialise our prime cache */
	atkinSieve(primeTarget)

	/* iterate along desks figuring out the number of presents they get */
	for ; currentMax < limit; desk+=step {
		presentsForDesk = calculatePresentsFromCache(desk)
		if presentsForDesk > currentMax {
			currentMax = presentsForDesk
		}
	}

	desk -= step

	// quit timing now we have the result
	duration = time.Since(start)

	fmt.Printf("%v\n", desk)
	fmt.Printf("%v\n", duration)
}
