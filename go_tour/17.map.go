package main

import (
	"fmt"
)

func main(){
	var m map[string]string
	m= make(map[string]string)


	m["name"]="shubham"
	fmt.Println(m)

	for entity:= range m{
		fmt.Println(entity)
	}

	for key,value:= range m{
		fmt.Println(key,value)
	}


}