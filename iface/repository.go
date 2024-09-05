package iface

import (
	"context"

	"github.com/shivdaskadam/golang-boilerplate/request_response/healthCheck"
	"github.com/shivdaskadam/golang-boilerplate/request_response/user"
	"github.com/shivdaskadam/golang-boilerplate/schemas"
)

type Repository interface {
	GetHealthCheck(ctx context.Context) (res healthCheck.Response, err error)
	GetUser(ctx context.Context) (res user.GetUserResponse, err error)
	CreateEmployee(employee *schemas.Employee) error
	GetEmployeeByID(id int) (*schemas.Employee, error)
	GetAllEmployees() ([]schemas.Employee, error)
	UpdateEmployee(employee *schemas.Employee) error
	DeleteEmployee(id int) error
}
