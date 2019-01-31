package main

import "fmt"

type Vertex struct {
	X int64
	Y int64
}
func main(){

	var m = map[string]Vertex{
		"house":Vertex{1,4},
		"school":Vertex{23,11},
	}
	fmt.Println(m)

	m["Resto"]=Vertex{112,43}
	fmt.Println(m)
}