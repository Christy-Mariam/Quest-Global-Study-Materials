package main

import "fmt"

func num(yPtr *int){
	*yPtr=5
}

func main(){
	yPtr := new(int)
	num(yPtr)
	fmt.Println(*yPtr)
}
