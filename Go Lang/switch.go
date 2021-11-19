package main

import "fmt"

func main(){
	length := 9
	switch{
	case length <=7:
		fmt.Println("Short")
	case length <=8:
		fmt.Println("Normal")
	case length >8:
		fmt.Println("Long")
	}
}
