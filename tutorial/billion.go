// package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("GO count to 1 Billion!")
	start := time.Now()
	for x := 1; x <= 1000000000; x++ {
	}
	elapsed := time.Since(start)
	log.Printf("Time taken is: %s", elapsed)
	fmt.Println("GO is even more amazing!, 12 lines of codes")
}