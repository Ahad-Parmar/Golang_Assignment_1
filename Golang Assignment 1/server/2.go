package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now()
	fooditem := []string{"tandoori chicken", "seekh paratha", "khapsa", "biryani"}
	results := make(chan bool)
	
	for _, f := range fooditem {
		go func(f string) {
			make(f)
			results <- true
		}(f)
	}
	for i := 0; i < len(fooditem); i++ {
		<-results
	}
	fmt.Printf("made in %v\n", time.Since(started))
}

func make(food string) {
	fmt.Printf("cooking %s...\n", food)
	time.Sleep(3 * time.Second)
	fmt.Printf("done making %s\n", food)
	fmt.Println("")
}
