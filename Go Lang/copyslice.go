package main

import "fmt"

func main(){
	slice1 := []int{8,9,10}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
