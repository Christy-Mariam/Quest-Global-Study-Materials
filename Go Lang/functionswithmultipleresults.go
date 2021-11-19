package main

import "fmt"

func mult(a, b string)(string, string){
	return a,b
}

func main(){
	x,y := mult("hello", "world")
	fmt.Println(x,y)
}
