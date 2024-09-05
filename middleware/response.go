package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GenerateCSRFToken generates a random CSRF token
func generateCSRFToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// ResponseTransformerMiddleware intercepts the response and transforms it
func ResponseTransformerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Continue with the request
		err := c.Next()
		if err != nil {
			return err
		}
		csrfToken, err := generateCSRFToken()
		if err != nil {
			return err
		}
		// Set security headers
		c.Cookie(&fiber.Cookie{
			Name:     "XSRF-TOKEN",
			Value:    csrfToken,
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Lax",
		})
		c.Set("Content-Security-Policy", "default-src 'self';")
		c.Set("Strict-Transport-Security", "max-age=32140800; includeSubDomains; preload")
		c.Set("Cache-control", "no-cache, no-store, must-revalidate")
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Frame-Options", "deny")
		c.Set("Pragma", "no-cache")
		c.Set("X-Permitted-Cross-Domain-Policies", "master-only")
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Platform, Access-Token")

		responseBody := c.Response().Body()
		var parsedBody interface{}

		if len(responseBody) > 0 {
			err = json.Unmarshal(responseBody, &parsedBody)
			if err != nil {
				log.Printf("Error unmarshaling response body: %v", err)
				parsedBody = string(responseBody)
			}
		}

		data := map[string]interface{}{
			"success":    true,
			"identifier": uuid.New().String(),
			"data":       parsedBody,
		}

		return c.JSON(data)
	}
}
