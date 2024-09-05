package iface

import (
	"context"

	"github.com/shivdaskadam/golang-boilerplate/request_response/healthCheck"
	"github.com/shivdaskadam/golang-boilerplate/request_response/user"
	"github.com/shivdaskadam/golang-boilerplate/schemas"
)

type Service interface {
	HealthCheck(ctx context.Context) (res healthCheck.Response, err error)
	GetUser(ctx context.Context) (res user.GetUserResponse, err error)
	CreateEmployee(ctx context.Context, employee *schemas.Employee) error
	GetEmployee(ctx context.Context, id int) (*schemas.Employee, error)
	GetEmployees(ctx context.Context) ([]schemas.Employee, error)
	UpdateEmployee(ctx context.Context, employee *schemas.Employee) error
	DeleteEmployee(ctx context.Context, id int) error
}
