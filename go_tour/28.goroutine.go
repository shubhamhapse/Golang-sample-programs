package main

import (
	"fmt"
	"time"
)

func sayWithWait(s string) {
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s,i)
	}
}
func sayWithoutWait(s string) {
	for i := 0; i < 100; i++ {
		fmt.Println(s,i)
	}
}
func main() {
	go sayWithoutWait("world")
	sayWithoutWait("hello")
}
