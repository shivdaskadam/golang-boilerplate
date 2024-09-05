package services

import (
	"context"
	"time"

	"github.com/shivdaskadam/golang-boilerplate/request_response/healthCheck"
)

func (s *service) HealthCheck(ctx context.Context) (res healthCheck.Response, err error) {

	res = healthCheck.Response{
		Message:   "Service is up and running",
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Simulate some processing time
	time.Sleep(1 * time.Second)
	return res, nil
}
