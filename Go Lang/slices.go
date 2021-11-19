package main

import "fmt"

func main(){
	slice := []int{7,9,10,12}
	fmt.Println("slice ==",slice)
	for i:=1; i<len(slice); i++{
		fmt.Printf("slice[%d] == %d\n", i, slice[i])
	}
}
