package main

import (
	"fmt"
	"reflect"
)

func main()  {
	 a :=[]int64{2,2,3,4,5}

	 for i:=range a{
	 	fmt.Println(reflect.TypeOf(i))
	 }
}