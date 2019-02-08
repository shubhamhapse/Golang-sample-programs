package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedInt int

var sc sync.Mutex
func increseInt(){
	for{
		sc.Lock()
		sharedInt++
		sc.Unlock()
		time.Sleep(time.Millisecond*100)
	}
}

func main(){
	i:=0
	for i< 5 {
		go increseInt()
		i++
	}
	time.Sleep(time.Second*10)
	sc.Lock()
	fmt.Println(sharedInt)
	sc.Unlock()
}