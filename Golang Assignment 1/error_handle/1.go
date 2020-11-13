package main

import (
	"errors"
	"fmt"
)

func main() {
	
	total, error := sumOf(5, 5)
	if error != nil {
		fmt.Println("error!!")
	} else {
		fmt.Println(total)
	}
}

func sf(start int, end int) (int, error) {
	if start > end {
		return 0, errors.New("start > end")
	}
	total := 0
	for i := start; i <= end; i++ {
		total += i
	}
	return total, nil
}