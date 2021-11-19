package main

import (
	"fmt"
	"math"
)

type Root struct {
	A, B float64
}

func (r *Root) Hyp() float64 {
	return math.Sqrt(r.A*r.A + r.B*r.B)
}

func main(){
	r := Root{5,6}
	fmt.Println("C =", r.Hyp())
}
