package main

import (
	"fmt"
)

func sum(a []int, ch chan int){
	sum := 0
	for _, v := range a {
		sum += v
	}
	ch <- sum
}
	
func main(){
	a := []int{7,0,-3,5,0,4}
	ch := make(chan int)
	go sum(a[:len(a)/2], ch)
	go sum(a[len(a)/2:], ch)
	x,y := <-ch, <-ch
	fmt.Println(x, y, x+y)
}


