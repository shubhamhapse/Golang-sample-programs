package main

import (
	"fmt"
	"time"
)

func sender(c chan int){
	a:=1
	for {
		c<-a
		fmt.Println("added new")
		a++
		time.Sleep(time.Second)
	}
}

func listener(c chan int){
	time.Sleep(time.Second*10)

	for {
		fmt.Println(<-c)
	}
}

func main(){

	//buffered channel
	c:= make(chan int,5)
	go listener(c)
	go sender(c)
	time.Sleep(time.Minute)
}