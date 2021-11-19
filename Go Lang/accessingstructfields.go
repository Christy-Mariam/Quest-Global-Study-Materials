package main

import "fmt"

type Rect struct{
	Width int
	Height int
}

func main(){
	r := Rect{1,2}
	r.Width = 18
	fmt.Println(r.Width)
}
