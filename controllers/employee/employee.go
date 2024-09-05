package employee

import (
	"strconv"

	"github.com/shivdaskadam/golang-boilerplate/iface"
	"github.com/shivdaskadam/golang-boilerplate/schemas"
	"github.com/gofiber/fiber/v2"
)

func CreateEmployeeHandler(svc iface.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		employee := new(schemas.Employee)
		if err := c.BodyParser(employee); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		err := svc.CreateEmployee(c.Context(), employee)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(employee)
	}
}

func GetEmployeeHandler(svc iface.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		intID, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID format",
			})
		}

		employee, err := svc.GetEmployee(c.Context(), intID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(employee)
	}
}

func GetEmployeesHandler(svc iface.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		employees, err := svc.GetEmployees(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(employees)
	}
}

func UpdateEmployeeHandler(svc iface.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid employee ID",
			})
		}
		employee := new(schemas.Employee)
		if err := c.BodyParser(employee); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		employee.Id = id
		err = svc.UpdateEmployee(c.Context(), employee)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(employee)
	}
}

func DeleteEmployeeHandler(svc iface.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		intID, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID format",
			})
		}
		err = svc.DeleteEmployee(c.Context(), intID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}
