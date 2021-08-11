package client

import "net/http"

type Client interface {
	Get(url string) (*http.Response, error)
}

type client struct {

}

func NewClient() Client {
	return client{}
}

func (c client) Get(url string) (*http.Response, error) {
	return http.Get(url)
}
