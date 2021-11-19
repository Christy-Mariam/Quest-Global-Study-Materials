package main

import (
	"fmt"
	"time"
)

func main(){
	go msg()
	fmt.Println("\nmsg from func main, I'm finished!")
	time.Sleep(time.Millisecond * 7500)
}

func msg(){
	for i := 1; i<=5; i++ {
		time.Sleep(time.Millisecond * 1000)
		if i >3 {
			fmt.Println(i, "seconds...yawn")
		}else{
			fmt.Println(i, "seconds")
		}
	}
	fmt.Println("\nmsg from func msg : I'm finished!")
}


