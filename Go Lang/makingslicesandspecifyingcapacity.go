package main

import "fmt"

func main(){
	a := make([]int, 4)
	printS1 ("a", a)
	b := make([]int, 0,4)
	printS1 ("b", b)
	c := b[:1]
	printS1 ("c", c)
	d := c[2:4]
	printS1 ("d", d)
}

func printS1(s string, x []int){
	fmt.Printf("%s len=%d cap=%d %v\n",s, len(x), cap(x), x)
}
