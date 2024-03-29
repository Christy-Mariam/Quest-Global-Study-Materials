package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
	"bytes"
)

func main(){
	response, err:= http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err != nil{
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}else{
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)
	request, _ := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err = client.Do(request)
	//response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil{
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}else{
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}