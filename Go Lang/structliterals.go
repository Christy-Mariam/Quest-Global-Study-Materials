package main

import "fmt"

type Rect struct{
	Width, Height int
}

var(
	r1= Rect{7,8}      //type Rect
	r2= Rect{Width:4}  //Height is implicitly 0
	r3= Rect{}         //Width and Height both 0
	p= &Rect{7,8}      //type * Rect
)

func main(){
	fmt.Println(r1, r2, r3, p)
}
