package main

import "fmt"

func calculateSum(functionVariable func(x int,y int) int,x,y int)int {
	return functionVariable(x,y)

}

func main() {
	sum:= func (x int,y int )int {
		return x+y
	}

	diff:= func(x,y int) int{
		return x-y
	}
	fmt.Println(calculateSum(sum,3,5))
	fmt.Println(calculateSum(diff,5,4))
}
