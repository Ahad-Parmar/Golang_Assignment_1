package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	started := time.Now()
	fooditem := []string{"tandoori chicken", "seekh paratha", "khapsa", "biryani"}
	var ap sync.WaitGroup
	wg.Add(len(fooditem))
	for _, f := range fooditem {
		go func(f string) {
			make(f)
			ap.Done()
		}(f)
	}
	ap.Wait()
	fmt.Printf("made in %v\n", time.Since(started))
}

func make(food string) {
	fmt.Printf("making in  %s...\n", food)
	time.Sleep(3 * time.Second)
	fmt.Printf("done %s\n", food)
	fmt.Println("")
}