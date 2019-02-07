package main

import (
	"fmt"
	"strings"
)

func main(){
	v:="my name is shubham my brother's name is also shubham"
	fmt.Println(countWords(v))

}

func countWords(input string) map[string]int64{
	wordlist:= strings.Split(input," ")
	mapResult:= map[string]int64{}
	for _,word :=range wordlist{
		_, status:= mapResult[word]
		if status{
			mapResult[word]=mapResult[word]+1
		}else {
			mapResult[word]=1
		}
	}
	return mapResult
}