package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// invoke server & receive response
	response, err := http.Get("http://localhost:8080/kello")
	if err != nil {
		fmt.Println("Failed to make an HTTP request:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return
	}

	fmt.Println("Response from server:", string(body))
}
