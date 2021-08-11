package client

import (
	"fmt"
	"gobreaker"
	"net/http"
)

type Client interface {
	Get(url string) (*http.Response, error)
}

type client struct {
	circuitBreaker *gobreaker.CircuitBreaker
}

func NewClient(circuitBreaker *gobreaker.CircuitBreaker) Client {
	return client{circuitBreaker}
}

func (c client) Get(url string) (*http.Response, error) {
	fmt.Println("Circuit breaker Name: ", c.circuitBreaker.Name())
	fmt.Println("Circuit breaker State: ", c.circuitBreaker.State())

	response, err := c.circuitBreaker.Execute(func() (interface{}, error) {
		return http.Get(url)
	})
	return response.(*http.Response), err
}
