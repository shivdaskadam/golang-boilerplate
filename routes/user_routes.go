package routes

import (
	"github.com/shivdaskadam/golang-boilerplate/controllers/user"
	svc "github.com/shivdaskadam/golang-boilerplate/iface"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterUserRoutes(app *fiber.App, service svc.Service, db *gorm.DB) {
	userGroup := app.Group("/user")
	userGroup.Get("/", user.GetUserHandler(service))
}
