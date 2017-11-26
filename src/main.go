package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
	"math"
)

type deskData struct {
	desk, presents uint32
}

type primeMultipleData struct {
	multiple, sum uint32 // n^x, sum of n^x + n^x-1 + ... n^0 (* 10)
}

type primeData struct {
	prime uint32
	multiples []primeMultipleData
}

var limit uint32

// here we assume we only need primes up to 65535; i.e. we'll deal only with numbers <= 4294836225 (almost all unsigned 32 bit integers)
var primes []primeData

func generateExponents(n uint32) []primeMultipleData {
	var i, cumulativeExponent, nPower uint32
	var primeExponentList []primeMultipleData
	cumulativeExponent = 0
	primeExponentList = append(primeExponentList, primeMultipleData{1, 10})
	primeExponentList = append(primeExponentList, primeMultipleData{n, n * 10 + 10})

	for i = 2; cumulativeExponent < limit; i++ {
		nPower = uint32((primeExponentList)[i-1].multiple * n)
		cumulativeExponent = uint32(primeExponentList[i-1].sum) + (nPower * 10)
		primeExponentList = append(primeExponentList, primeMultipleData{nPower, cumulativeExponent})
	}

	return primeExponentList
}

func atkinSieve(max uint32) {
	var x, y, n uint32
	maxSqrt := math.Sqrt(float64(max))

	isPrime := make([]bool, max, max)

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

func deskDelivery(deskDataChannel chan deskData, quitChannel chan bool) {
	deskData := deskData{0,0}
	for {
		select {
		case deskDataChannel <- deskData:
			deskData.desk += 12
			deskData.presents = calculatePresents(deskData.desk)
		case <-quitChannel:
			return
		}
	}
}


func main() {
	var desk, target, primeTarget uint32
	var start time.Time
	var duration time.Duration
	//var deskData deskData
	//deskDataChannel := make(chan deskData)
	//quitChannel := make(chan bool)

	parsedTarget, _ := strconv.ParseInt(os.Getenv("PRESENTS"), 10, 64)
	target = uint32(parsedTarget)
	limit = target
	primeTarget = uint32(math.Sqrt(float64(target)))

	start = time.Now()

	//go func() {
	//	for deskData = range deskDataChannel {
	//		if deskData.presents > currentMax {
	//			fmt.Printf("%v: %v\n", deskData.desk, deskData.presents)
	//			currentMax = deskData.presents
	//			desk = deskData.desk
	//		}
	//		if deskData.presents > target {
	//			quitChannel <- true
	//		}
	//	}
	//}()
	//
	//deskDelivery(deskDataChannel, quitChannel)

	//for i = 0; currentMax < target; i+=12 {
	//	presentsForDesk = calculatePresents(i)
	//	if presentsForDesk > currentMax {
	//		fmt.Printf("%v: %v\n", i, presentsForDesk)
	//		currentMax = presentsForDesk
	//		desk = i
	//	}
	//}

	atkinSieve(primeTarget)

	duration = time.Since(start)

	fmt.Printf("%v\n", primes)

	fmt.Printf("%v\n", desk)
	fmt.Printf("%v\n", duration)
}
