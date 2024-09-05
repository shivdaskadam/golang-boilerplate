package repository

import (
	"github.com/shivdaskadam/golang-boilerplate/schemas"
)

func NewEmployeeRepository(repo *repository) *repository {
	return &repository{repo.db}
}

func (repo *repository) CreateEmployee(employee *schemas.Employee) error {
	return repo.DB().Create(employee).Error
}

func (repo *repository) GetEmployeeByID(id int) (*schemas.Employee, error) {
	var employee schemas.Employee
	err := repo.DB().First(&employee, "id = ?", id).Error
	return &employee, err
}

func (repo *repository) GetAllEmployees() ([]schemas.Employee, error) {
	var employees []schemas.Employee
	err := repo.DB().Find(&employees).Error
	return employees, err
}

func (repo *repository) UpdateEmployee(employee *schemas.Employee) error {
	return repo.DB().Save(employee).Error
}

func (repo *repository) DeleteEmployee(id int) error {
	return repo.DB().Delete(&schemas.Employee{}, "id = ?", id).Error
}
