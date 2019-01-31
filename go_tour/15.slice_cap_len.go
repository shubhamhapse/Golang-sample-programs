package main

import "fmt"

func main(){
	var a = []int{1,2,3,4,5}
	fmt.Println("cap",cap(a),"len",len(a) )
	a= append(a,6)
	fmt.Println("cap",cap(a),"len",len(a) )


	var b = []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("cap",cap(b),"len",len(b) )
	b= append(b,6)
	fmt.Println("cap",cap(b),"len",len(b) )

}