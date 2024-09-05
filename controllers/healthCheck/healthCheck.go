package healthCheck

import (
	"github.com/shivdaskadam/golang-boilerplate/iface"
	"github.com/gofiber/fiber/v2"
)

func HealthCheckHandler(svc iface.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Call the HealthCheck method from the service
		res, err := svc.HealthCheck(c.Context())
		if err != nil {
			// Return a JSON error response using Fiber's Status and JSON methods
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Return a JSON response with the result
		return c.JSON(res)
	}
}
