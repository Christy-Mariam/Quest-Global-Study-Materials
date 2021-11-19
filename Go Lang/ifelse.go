package main

import "fmt"

func main(){
	for z:=1; z<=5; z++{
		if(z % 2==0){
		fmt.Println(z,"even number")
	}else{
		fmt.Println(z,"odd number")
		}
	}
}
