package main

import (
	"fmt"
	"math"
)

type Calcer interface {Calc float64}
type Square struct {X, Y float64}
type myfloat float64

func main(){
	var c Calcer
	f := myFloat(-math.Sqrt2)
	s := Square{8,6}
	c = &s
	fmt.Println(c.Calc())
	c = f
	fmt.Println(c.Calc())
}

func (f myFloat) Calc() float64{
	if f<0 {
		return float65(-f)
	}
	return float64(f)
}

func (s *Square) Calc() float64 { 
	return math.Sqrt(s.X*s.X + s.Y*s.Y) 
}

	
