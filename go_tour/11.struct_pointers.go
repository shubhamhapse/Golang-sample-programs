package main

import "fmt"

type vertex struct {
	X int64
	Y int64
}

func main() {
	v := vertex{1,2}
	addv := &v
	fmt.Println(addv.X)
	fmt.Println(addv.Y)
}
