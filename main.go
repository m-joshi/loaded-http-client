package main

import (
	"fmt"
	"loaded-http-client/client"
)

func main() {
	fmt.Println("Starting loaded http client")

	loadedClient := client.NewClient()
	response, err := loadedClient.Get("http://localhost:8080/simpleServer")
	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println("RESPONSE: ", response)
	}
}
