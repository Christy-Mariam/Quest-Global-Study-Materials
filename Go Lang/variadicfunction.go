package main

import "fmt"

func add(args ...int)int{
	total := 0
	for _, n:= range args{
		total += n
	}
	return total
}

func main(){
	fmt.Println(add(8,13,4,6,3))
}
