package main

import "fmt"

type vertex struct {
	x int64
	y int64
}

type vertexExp struct {
	X int64
	Y int64
}

type vertexNotExp struct {
	x int64
	y int64
}
func main() {
	s := [2]vertex{
		vertex{1, 2},
		vertex{-1, 12},
	}
	fmt.Println(s)
	fmt.Println(s[0].x, s[0].y)

	sExp:= vertexExp{1,2}
	fmt.Println(sExp.X, sExp.Y)

	sNotExp := vertexNotExp{1,2}
	fmt.Println(sNotExp.x, sNotExp.y)
//conclusion: Export is applicable only for differnt package
	}

