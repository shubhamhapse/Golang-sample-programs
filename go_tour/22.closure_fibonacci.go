package main

import "fmt"

func fibonacci() func () int{
	pre, next:=-1,1

	return func() int{
		temp:=pre+next
		pre=next
		next=temp
		return temp
	}
}

func main(){
	fib:=fibonacci()
	for i:=0;i<10;i++{
		fmt.Print(fib())
		fmt.Print(" ")
	}
}