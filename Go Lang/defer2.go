package main

import "fmt"

func main(){
		fmt.Println("\nLast In---->First Out\n")
		for z:=1; z<5; z++{
			defer fmt.Println(z, "popped")
			fmt.Println(z, "deferred")
		}
		fmt.Println("\n")
}
