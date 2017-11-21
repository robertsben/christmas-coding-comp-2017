package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("Hello, world\n")
	execTime := time.Since(start)
	fmt.Printf("%v\n", execTime)
}
