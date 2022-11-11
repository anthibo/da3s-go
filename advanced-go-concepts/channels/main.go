package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}
func main() {
	// buffered channel
	// dataChan := make(chan int, 2)

	dataChan := make(chan int)

	// dataChan <- 789 // add data to channel
	// n := <-dataChan // receive data from the channel
	// fmt.Printf("n = %d\n", n)

	// runs in background goroutine
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()
	// consume data from the channel
	for n := range dataChan {
		fmt.Printf("n = %d\n", n)
	}

}
