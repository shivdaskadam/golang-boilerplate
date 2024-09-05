package internal

import (
	"fmt"

	svc "github.com/shivdaskadam/golang-boilerplate/iface"
	implementation "github.com/shivdaskadam/golang-boilerplate/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/shivdaskadam/golang-boilerplate/repository"
	"github.com/shivdaskadam/golang-boilerplate/routes"
)

func StartApp(db *gorm.DB, app *fiber.App) {
	var repo svc.Repository
	{
		var err error
		repo, err = repository.New(db)
		if err != nil {
			fmt.Println("error occurred while creating repo instance")
		}
	}
	var service svc.Service
	{
		service = implementation.NewService(repo)
	}

	routes.RegisterRoutes(app, service, db)
}
