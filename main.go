package main

import (
	"fmt"
	circuit_breaker "loaded-http-client/circuit-breaker"
	"loaded-http-client/client"
)

var loadedClient client.Client

func init() {
	fmt.Println("Initialising loaded http client")
	fmt.Println(".")
	fmt.Println("..")
	fmt.Println("...")

	circuitBreakerFactory := circuit_breaker.NewFactory()
	circuitBreaker := circuitBreakerFactory.CircuitBreaker("Client circuit breaker", 2, 60, 15)
	loadedClient = client.NewClient(circuitBreaker)
}

func main() {
	fmt.Println("Starting loaded http client")
	fmt.Println(".")
	fmt.Println("..")
	fmt.Println("...")

	response, err := loadedClient.Get("http://localhost:8080/simpleServer")
	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println("RESPONSE: ", response)
	}
}
