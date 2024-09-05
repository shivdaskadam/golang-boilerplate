package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/shivdaskadam/golang-boilerplate/request_response/healthCheck"
)

func (repo *repository) GetHealthCheck(ctx context.Context) (res healthCheck.Response, err error) {
	res.Timestamp = fmt.Sprintf("%d", time.Now().UTC().Unix())
	res.Message = "Service is up and running"
	return res, nil
}
