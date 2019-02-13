package main

import "fmt"

const (
	a=iota
	b=iota
	c
	d
)

const e  = iota

const f,g,h  = iota,iota,iota //same line has same value

func main(){
	fmt.Println(a,b,c,d)
	fmt.Println(e)
	fmt.Println(f,g,h)
}