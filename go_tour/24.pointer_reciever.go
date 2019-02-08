package main

import "fmt"

type vertex struct {
	X,Y int64
}

func (v *vertex) shiftPoint(){
	v.X=v.X+5
	v.Y=v.Y+5
}

func (v vertex) shiftPointnative(){
	v.X=v.X+5
	v.X=v.X+5
}


func main(){
	v := vertex{1,3}
	v.shiftPointnative()
	fmt.Println("without pointer",v)
	v.shiftPoint()
	fmt.Println("with pointer",v)

	//methods that receives pointer receiver can be called on pointer or variable
	vPointer:=&v
	vPointer.shiftPoint()
	fmt.Println(*vPointer)
	vPointer.shiftPointnative()
	fmt.Println(vPointer)
}