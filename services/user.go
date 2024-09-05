package services

import (
	"context"

	"github.com/shivdaskadam/golang-boilerplate/request_response/user"
)

func (s *service) GetUser(ctx context.Context) (res user.GetUserResponse, err error) {
	res, err = s.repository.GetUser(ctx)
	return res, err
}
