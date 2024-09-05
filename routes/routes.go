package routes

import (
	svc "github.com/shivdaskadam/golang-boilerplate/iface"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, service svc.Service, db *gorm.DB) {
	RegisterEmployeeRoutes(app, service, db)
	RegisterHealthCheckRoutes(app, service, db)
	RegisterUserRoutes(app, service, db)
}
