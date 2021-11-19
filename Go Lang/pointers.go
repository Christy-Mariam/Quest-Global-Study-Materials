package main

import "fmt"

func fun(yPtr *int){
	*yPtr=0
}

func main(){
	y := 10
	fun (&y)
	fmt.Println(y)
}
