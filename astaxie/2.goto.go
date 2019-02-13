package main

import "fmt"

func main()  {
	i:=0
	here:
		fmt.Println(i)
		i++
		if i<10{
			goto here
		}
}