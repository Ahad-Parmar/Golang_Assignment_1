package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var total = 1000
	ch := make(chan int, total)
	single := make([]int, total, total)
	fmt.Printf("random number generation.\n")
	start := time.Now()
	go MRN(total, ch)  // MRN- make Random Numbers
	for i := 0; i < total; i++ {
		single = append(single, (<-ch))
	}
	ESR := time.Since(start)  //elapsed Single Run
	fmt.Printf("Run took %s\n", ESR)
}

func MRN(numInts int, ch chan int) {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	for i := 0; i < numInts; i++ {
		ch <- generator.Intn(numInts)
	}
	close(ch)

	for item := range ch {
		fmt.Println(item)
	}
}
