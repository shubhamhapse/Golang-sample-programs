package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X int64
	Y int64
}

func (v Vertex) distFromCenter() float64{

	return math.Sqrt(float64(v.X*v.X+v.Y*v.Y))
}

func main()  {
	v:=Vertex{4,3}
	fmt.Println(v.distFromCenter())
	fmt.Println(v)

}