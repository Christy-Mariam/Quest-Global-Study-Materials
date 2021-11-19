package main

import "fmt"

func main(){
	slice1 := []int{8,9,10}
	slice2 := append(slice1, 11, 12)
	fmt.Println(slice1, slice2)
}
