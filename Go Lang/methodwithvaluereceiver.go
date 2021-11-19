package main

import (
	"fmt"
	"math"
)

type Root float64
	func (r Root) Abs() float64 {
	if r<0 {
		return float64(-r)
	}
	return float64(r)
}

func main(){
	r := Root(-math.Sqrt2)
	fmt.Println(r.Abs())
}
