package main

import "fmt"

var valueMap map[string]string

func main(){
	valueMap= make(map[string]string)
	mainMap:=make(map[int64]map[string]string)
	valueMap["name"]="shubham"
	valueMap["address"]="Pune"
	mainMap[1]=valueMap

	fmt.Println(mainMap)

	mainMap[2]= map[string]string{"name":"sam","addrss":"Sangli"}

	fmt.Println(mainMap)


}