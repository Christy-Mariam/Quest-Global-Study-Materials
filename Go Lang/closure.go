package main

import "fmt"

func main(){
	a := 4
	decrement := func()int{
		a--
		return a
	}
	fmt.Println(decrement())
	fmt.Println(decrement())
}
