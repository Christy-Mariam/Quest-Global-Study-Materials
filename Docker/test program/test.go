package main

import(
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",SimpleServer)
	http.ListenAndServe(":9055", nil)
}

func SimpleServer(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World!")
}