package middleware

import (
	"bytes"
	"context"
	"log"
	"time"

	"github.com/shivdaskadam/golang-boilerplate/config"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type RequestLogger struct {
	collection *mongo.Collection
}

func NewRequestLogger(client *mongo.Client) *RequestLogger {
	collection := config.GetCollection("service_logger")
	// collection := client.Database("core-bsa").Collection("service_logger")
	return &RequestLogger{collection: collection}
}

func (r *RequestLogger) LogRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Read the request body
		reqBody := c.Body()
		reqBodyStr := string(reqBody)

		// Copy the raw request body to log it later
		rawReq := string(c.Request().Header.Header()) + "\n" + reqBodyStr

		// Capture the response body
		var resBody bytes.Buffer
		c.Response().SetBodyStream(&resBody, -1)

		// Execute the next handler in the chain
		err := c.Next()

		// Read the response body
		responseBody := c.Response().Body()

		// Log entry
		logEntry := map[string]interface{}{
			"requestID":      uuid.New().String(),
			"method":         c.Method(),
			"url":            c.OriginalURL(),
			"query_params":   c.OriginalURL(),
			"request_params": c.Params("*"),
			"request_body":   reqBodyStr,
			"raw_request":    rawReq,
			"response":       string(responseBody),
			"status":         c.Response().StatusCode(),
			"latency":        time.Since(start).String(),
			"user_agent":     c.Get("User-Agent"),
			"created_at":     time.Now().Format(time.RFC3339),
		}

		_, err = r.collection.InsertOne(context.Background(), logEntry)
		if err != nil {
			log.Printf("Failed to log request: %v", err)
		}

		return err
	}
}
