package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomnum int
}
type Result struct {
	job Job
	sod int
}

var jobs = make(chan Job, 5)
var results = make(chan Result, 5)

func digits(number int) int {
	sum := 0
	num := number
	for num != 0 {
		digit := no % 10
		sum += digit
		num /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomnum)}
		results <- output
	}
	wg.Done()
}
func createWorkerPool(noofworkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noofworkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}
func allocate(noofjobs int) {
	for i := 0; i < noofjobs; i++ {
		randomnum := rand.Intn(999)
		job := Job{i, randomnum}
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomnum, result.sod)
	}
	done <- true
}
func main() {
	startTime := time.Now()
	noofjobs := 50
	go allocate(noofjobs)
	done := make(chan bool)
	go result(done)
	noofworkers := 10
	createWorkerPool(noofworkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time ", diff.Seconds(), "seconds")
}