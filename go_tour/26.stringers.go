package main

import "fmt"

type Person struct {
	Name string
	Address string
}

func (p Person) String() string{
	return "Persons name: "+p.Name+ " Address: "+p.Address
}

func main()  {
	p:= Person{"Shubham","Baner"}
	var i interface{}=p
	fmt.Println(i)
	fmt.Println(p)
	//above both will use String() that is overrided
}