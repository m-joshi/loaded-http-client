package circuit_breaker

import (
	"gobreaker"
	"time"
)

type Factory interface {
	CircuitBreaker(name string, maxRequests uint32, interval time.Duration, timeout time.Duration) *gobreaker.CircuitBreaker
}

type factory struct {

}

func NewFactory() Factory {
	return factory{}
}

func (f factory) CircuitBreaker(name string, maxRequests uint32, interval time.Duration, timeout time.Duration) *gobreaker.CircuitBreaker {
	cbSettings := gobreaker.Settings{
		Name:        name,
		MaxRequests: maxRequests,
		Interval:    interval,
		Timeout:     timeout,
	}

	return gobreaker.NewCircuitBreaker(cbSettings)
}
