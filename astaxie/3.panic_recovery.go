package main

import "fmt"

func createPanic(){
	fmt.Println("Paniking")
	panic("Dude something is wrong")
}
func recoverPanic(){
	if r:=recover(); r!=nil{
		fmt.Println("Found panic")
		fmt.Println("Recovered")
	}else{
		fmt.Println("no panic found")
	}
}
func main(){
	fmt.Println("Started main thread")
	defer recoverPanic()
	createPanic()
	fmt.Println("Will not exexute this bcz of panic")


}