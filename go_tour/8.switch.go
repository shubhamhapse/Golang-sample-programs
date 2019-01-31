package main

import "fmt"

func main() {
	var no = int64(5)

	switch no {
	case 4:
		fmt.Println("4 is here")
	case 6:
		fmt.Println("6 is here")
	default:
		fmt.Println("default here")
	}
}
