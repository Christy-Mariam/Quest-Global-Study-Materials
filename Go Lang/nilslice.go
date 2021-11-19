package main

import "fmt"

func main(){
	var x [] int
	fmt.Println(x, len(x), cap(x))
	if x==nil{
		fmt.Println("nil!")
	}
}
