package main

import "fmt"

type I interface {
	printString()
}

type address struct {
	addr string
}

func (a address) printString() {
	fmt.Println(a.addr)
}

func main() {
	var i I
	addr := address{"asdfaf"}
	i = addr
	i.printString()
	i = &addr
	i.printString()
}
