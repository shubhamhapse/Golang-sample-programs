package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	totalThreads = 2
)

var wg sync.WaitGroup

func main() {
	wg.Add(totalThreads)

	go func() {
		printRandom(1)
		wg.Done()
	}()

	go func() {
		printRandom(2)
		wg.Done()
	}()
	wg.Wait()
}

func printRandom(delay time.Duration) {

	for count := 0; count < 10; count++ {
		fmt.Println(rand.Int())
		time.Sleep(time.Second * delay)
	}

}
