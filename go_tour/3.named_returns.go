package main

import "fmt"

func add(a,b int) (c int){
	c=a+b
	return
}

func main(){
	fmt.Println(add(5,6))
}