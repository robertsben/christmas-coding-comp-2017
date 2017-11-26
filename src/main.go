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
	var i, cumulativeExponent, nPower uint32
	var primeExponentList []primeMultipleData
	cumulativeExponent = 0
	primeExponentList = append(primeExponentList, primeMultipleData{1, 1})
	primeExponentList = append(primeExponentList, primeMultipleData{n, n + 1})

	for i = 2; cumulativeExponent <= limit; i++ {
		nPower = uint32(primeExponentList[i-1].multiple * n)
		cumulativeExponent = uint32(primeExponentList[i-1].sum) + nPower
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
	var total, i, j, primeTotal uint32
	total = 1
	for i = 0; i < uint32(len(primes)); i++ {
		primeTotal = 1
		for j = 1; desk % primes[i].multiples[j].multiple == 0 && j < uint32(len(primes[i].multiples)); j++ {
			primeTotal = primes[i].multiples[j].sum
		}
		total *= primeTotal
	}
	// return total of exponent sums, times the elf present delivery multiplier
	return total * 10
}


func main() {
	var desk, primeTarget, deskSearchStart, currentMax, presentsForDesk, step uint32
	var start time.Time
	var duration time.Duration
	step = 2
	//var deskData deskData
	//deskDataChannel := make(chan deskData)
	//quitChannel := make(chan bool)

	parsedTarget, _ := strconv.ParseInt(os.Getenv("PRESENTS"), 10, 64)

	start = time.Now()
	limit = uint32(parsedTarget)
	primeTarget = uint32(math.Sqrt(float64(limit)))

	/*
		We know that the desk will be greater than the root of the present number (or the actual number for values
		19 or below) so we set that as the start point for searching
	 */
	if primeTarget > 19 {
		deskSearchStart = primeTarget
	} else {
		deskSearchStart = uint32(math.Sqrt(float64(limit/10)))
	}

	if deskSearchStart % 2 != 0 {
		deskSearchStart--
	}

	/* initialise our prime cache */
	atkinSieve(primeTarget)

	for desk = deskSearchStart; currentMax < limit; desk+=step {
		presentsForDesk = calculatePresentsFromCache(desk)
		if presentsForDesk > currentMax {
			//fmt.Printf("%v, %v\n", desk, presentsForDesk)
			currentMax = presentsForDesk
		}
	}

	desk -= step

	duration = time.Since(start)

	fmt.Printf("%v\n", desk)
	fmt.Printf("%v\n", duration)
}
