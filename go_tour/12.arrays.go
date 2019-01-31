package main

import (
	"fmt"
)

func main()  {
	a:= [2]string{}
	a[0]="hello"
	a[1]="world"

	fmt.Println(a,a[0],a[1],&a)
}