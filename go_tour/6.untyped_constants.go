package main

import (
	"fmt"
)

const (
	Big = 1<<100
	Small = 1<<1
)

func itsInt(x int) int {
	return x
}

func itsFloat64(x float64) float64 {
	return x
}

func itsFloat32 (x float32) float32{
	return  x
}

func main(){
	//fmt.Println(itsInt(Big))
	fmt.Println(itsInt(Small))
	fmt.Println(itsFloat32(Small))
	fmt.Println(itsFloat32(Big))
	fmt.Println(itsFloat64(Small))
	fmt.Println(itsFloat64(Big))
}