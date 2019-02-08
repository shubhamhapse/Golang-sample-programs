package main

import (
	"fmt"
	"time"
)

type MyError struct {
	Time time.Time
	Details string
}
func (e MyError)Error() string{
	return e.Time.String()+e.Details
}

func createError() MyError{
	return MyError{time.Now(),"  Invalid I/P"}
}

type S string

func (s S)Error() string{
	return "Invalid i/p test"
}
func divideInt(a,b int) (int, error){
	if b==0{
		return 0, createError()
	}
	if b<0{
		return 0, S("asdfasdf")
	}
	return a/b,nil
}

func main(){
	a,b:=10,5
	fmt.Println(divideInt(a,b))
	a,b=10,0
	fmt.Println(divideInt(a,b))

	a,b=10,-4
	fmt.Println(divideInt(a,b))
}