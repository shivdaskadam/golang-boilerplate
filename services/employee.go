package services

import (
	"context"

	"github.com/shivdaskadam/golang-boilerplate/schemas"
)

func (s *service) CreateEmployee(ctx context.Context, employee *schemas.Employee) error {
	return s.repository.CreateEmployee(employee)
}

func (s *service) GetEmployee(ctx context.Context, id int) (*schemas.Employee, error) {
	return s.repository.GetEmployeeByID(id)
}

func (s *service) GetEmployees(ctx context.Context) ([]schemas.Employee, error) {
	return s.repository.GetAllEmployees()
}

func (s *service) UpdateEmployee(ctx context.Context, employee *schemas.Employee) error {
	return s.repository.UpdateEmployee(employee)
}

func (s *service) DeleteEmployee(ctx context.Context, id int) error {
	return s.repository.DeleteEmployee(id)
}
