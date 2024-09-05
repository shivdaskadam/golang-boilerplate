package routes

import (
	"github.com/shivdaskadam/golang-boilerplate/controllers/healthCheck"
	svc "github.com/shivdaskadam/golang-boilerplate/iface"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterHealthCheckRoutes(app *fiber.App, service svc.Service, db *gorm.DB) {
	healthCheckGroup := app.Group("/healthCheck")
	healthCheckGroup.Get("/", healthCheck.HealthCheckHandler(service))
}
