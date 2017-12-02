package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
	"math"
	"flag"
	"log"
	"runtime/pprof"
	"runtime"
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

var cpuprofile = flag.String("cpuprofile", fmt.Sprintf("cpu-%v.prof", time.Now()), "write cpu profile `file`")
var memprofile = flag.String("memprofile", fmt.Sprintf("mem-%v.prof", time.Now()), "write memory profile to `file`")

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

func calculatePrimeTotalForDesk(desk uint32, data primeData, totalChannel chan<- uint32) {
	var j uint16
	var total uint32

	for j = 1; j < uint16(len(data.multiples)) && desk % data.multiples[j].multiple == 0; j++ {
		total = data.multiples[j].sum
	}

	if total != 0 {
		totalChannel <- total
	}
}

func calculatePresentsFromCache(desk uint32) uint32 {
	var total, primeTotal uint32
	//var j uint16
	var prime primeData
	total = 1
	primeTotalsChannel := make(chan uint32)

	go func() {
		for primeTotal = range primeTotalsChannel {
			total *= primeTotal
		}
	}()

	for _, prime = range primes {
		go calculatePrimeTotalForDesk(desk, prime, primeTotalsChannel)
		//primeTotal = 1
		//for j = 1; j < uint16(len(prime.multiples)) && desk % prime.multiples[j].multiple == 0; j++ {
		//	primeTotal = prime.multiples[j].sum
		//}
		//total *= primeTotal
	}
	//for i = 0; i < uint32(len(primes)); i++ {
	//}
	// return total of exponent sums, times the elf present delivery multiplier
	return total * 10
}


func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	var desk, primeTarget, deskSearchStart, currentMax, presentsForDesk, step uint32
	var start time.Time
	var duration time.Duration
	step = 2
	//var deskData deskData
	//deskDataChannel := make(chan deskData)
	//quitChannel := make(chan bool)

	parsedTarget, _ := strconv.ParseInt(os.Getenv("PRESENTS"), 10, 64)

	// start timing now we have all the information
	start = time.Now()
	limit = uint32(parsedTarget)
	primeTarget = uint32(math.Sqrt(float64(limit)))
	fmt.Printf("%v\n", limit)

	/*
		Desk is always > limit/50 (as long as the result >= 2)
	 */
 	if limit >= 100 {
		deskSearchStart = limit/50
	} else {
		step = 1
		deskSearchStart = 1
	}

	if deskSearchStart > 1 && deskSearchStart % 2 != 0 {
		deskSearchStart--
	}

	/* initialise our prime cache */
	atkinSieve(primeTarget)

	/* iterate along desks figuring out the number of presents they get */
	for desk = deskSearchStart; currentMax < limit; desk+=step {
		presentsForDesk = calculatePresentsFromCache(desk)
		if presentsForDesk > currentMax {
			fmt.Printf("%v, %v\n", desk, presentsForDesk)
			currentMax = presentsForDesk
		}
	}

	desk -= step

	// quit timing now we have the result
	duration = time.Since(start)

	fmt.Printf("%v\n", desk)
	fmt.Printf("%v\n", duration)

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
