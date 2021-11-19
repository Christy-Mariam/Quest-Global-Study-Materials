package main

import "fmt"

func main(){
	slice := []int{7,9,10,12}
	fmt.Println("slice ==",slice)
	fmt.Println("slice[1:4] ==",slice[1:4])
	fmt.Println("slice[:3]  ==",slice[:3])
	fmt.Println("slice[2:]  ==",slice[2:])
}
