package main

import "fmt"

type Person struct {
	Name string
	Address string
}

type Student struct {
	Person
	CollegeName string
}

func (p Person) printPerson()  {
	fmt.Println(p.Name,p.Address)
}
func main()  {
	p:=Person{"Shubham","Hapse"}
	s:= Student{Person{"Ram","Hapse"},"walchand"}
	p.printPerson()
	s.printPerson()

	//field inheritance: all fields of anonymous struct are available in parent struct
	fmt.Println(s.Address,s.Name,s.CollegeName)
}