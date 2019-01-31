package main

import "fmt"

func main() {
	a := make([]int, 0, 14)
	fmt.Println("cap ",cap(a), " len ", len(a))
}
