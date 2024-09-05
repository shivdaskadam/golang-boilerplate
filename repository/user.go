package repository

import (
	"context"

	"github.com/shivdaskadam/golang-boilerplate/request_response/user"
)

func (repo *repository) GetUser(ctx context.Context) (res user.GetUserResponse, err error) {
	// DB operations
	// repo.Sql.GET()
	// repo.Mongo.Find()
	res = user.GetUserResponse{
		Id:   1,
		Name: "Kiran",
		Age:  25,
	}
	return res, nil
}
