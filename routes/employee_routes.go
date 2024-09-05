package routes

import (
	"github.com/shivdaskadam/golang-boilerplate/controllers/employee"
	svc "github.com/shivdaskadam/golang-boilerplate/iface"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterEmployeeRoutes(app *fiber.App, service svc.Service, db *gorm.DB) {
	employeeGroup := app.Group("/employee")
	employeeGroup.Post("/", employee.CreateEmployeeHandler(service))
	employeeGroup.Get("/:id", employee.GetEmployeeHandler(service))
	employeeGroup.Get("/", employee.GetEmployeesHandler(service))
	employeeGroup.Put("/:id", employee.UpdateEmployeeHandler(service))
	employeeGroup.Delete("/:id", employee.DeleteEmployeeHandler(service))
}
