package main

import (
	"fmt"
	"time"
)

func ping(ping <-chan int, pong chan<- int) {
	for {
		<-ping
		fmt.Println("ping")
		time.Sleep(time.Second)
		pong <- 1
	}
}

func pong(ping chan<- int, pong <-chan int) {
	for {
		<-pong
		fmt.Println("pong")
		time.Sleep(time.Second)
		ping <- 1
	}
}

func main() {
	pin := make(chan int)
	pon := make(chan int)

	go ping(pin, pon)
	go pong(pin, pon)

	pin <- 1 

	for {
		time.Sleep(time.Second)
	} 
}
