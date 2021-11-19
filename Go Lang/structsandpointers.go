package main

import "fmt"

type Rect struct{
	Width int
	Height int
}

func main(){
	r := Rect{4,6}
	p := &r
	p.Width = 18
	fmt.Println(r)
}
