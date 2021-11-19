package main

import "fmt"

func main(){
		fmt.Println("counting")
		for z:=1; z<5; z++{
			defer fmt.Println(z)
		}
		fmt.Println("finished")
}
