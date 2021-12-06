package main

import(
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",SimpleServer)
	http.ListenAndServe(":32100", nil)
}

func SimpleServer(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World!")
}