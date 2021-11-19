package main

import "fmt"

func f(msg string){
	for i := 0; i<10; i++ {
		fmt.Println(msg, ":", i)
	}
}

func main(){
	go f("value of i")
	var input string
	fmt.Scanln(&input)
}
