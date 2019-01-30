package main

import "fmt"

func checkEven(no int64) bool {
	if reminder:=no%2; reminder==0{
		return true
	}
	return false
}

func main(){
	fmt.Println(checkEven(4))
	fmt.Println(checkEven(5))

}