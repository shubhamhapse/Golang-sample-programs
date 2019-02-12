package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
func listen(c chan int) {
	for {
		i, ok := <-c;
		if ok {
			fmt.Println(i)
		} else {
			break
		}
	}
	wg.Done()
}

func sender(c chan int){
	for i:=0;i<10;i++{
		c<-i
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main()  {
	wg.Add(2)
	c:=make(chan int)
}
