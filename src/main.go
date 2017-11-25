package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
)

type DeskData struct {
	desk, presents uint32
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

func deskDelivery(deskDataChannel chan DeskData, quitChannel chan bool) {
	deskData := DeskData{0,0}
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
	var desk, target, currentMax uint32
	var start time.Time
	var duration time.Duration
	var deskData DeskData
	deskDataChannel := make(chan DeskData)
	quitChannel := make(chan bool)

	parsedTarget, _ := strconv.ParseInt(os.Getenv("PRESENTS"), 10, 64)
	target = uint32(parsedTarget)
	currentMax = 0

	start = time.Now()

	go func() {
		for deskData = range deskDataChannel {
			if deskData.presents > currentMax {
				fmt.Printf("%v: %v\n", deskData.desk, deskData.presents)
				currentMax = deskData.presents
				desk = deskData.desk
			}
			if deskData.presents > target {
				quitChannel <- true
			}
		}
	}()

	deskDelivery(deskDataChannel, quitChannel)

	//for i = 0; currentMax < target; i+=12 {
	//	presentsForDesk = calculatePresents(i)
	//	if presentsForDesk > currentMax {
	//		fmt.Printf("%v: %v\n", i, presentsForDesk)
	//		currentMax = presentsForDesk
	//		desk = i
	//	}
	//}

	duration = time.Since(start)

	fmt.Printf("%v\n", desk)
	fmt.Printf("%v\n", duration)
}
