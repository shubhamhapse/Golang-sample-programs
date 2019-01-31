package main

import "fmt"

func main() {
	var a []int
	a = append(a, 5)
	a = append(a, 4)
	a = append(a, 3)
	a = append(a, 2)
	a = append(a, 1)
	a = append(a, 0)
	fmt.Println(a)

	//delete index 3
	a = append(a[:3], a[4:]...)
	fmt.Println(a)
}
