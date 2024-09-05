package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/shivdaskadam/golang-boilerplate/config"
	"github.com/shivdaskadam/golang-boilerplate/internal"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// helmet
	app.Use(helmet.New())

	// Logging
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// CORS
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Custom-Handler", "hello")
		c.Set("Allow-Origins", "*")
		c.Set("Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Platform, Access-Token")
		return c.Next()
	})

	db, err := config.InitDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Connect to MongoDB
	client, cancel, err := config.InitMongo()
	if err != nil {
		log.Fatal(err)
	}
	defer config.DisconnectMongo(client, cancel)

	// Health API
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/health",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			// check for db/redis connections
			// Connect to MySQL db
			db.Exec("SELECT 1")
			config.PingDB(client)
			return true
		},
		ReadinessEndpoint: "/ready",
	}))

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorization, Platform, Access-Token",
	// }))

	// Recover
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// app.Use(middleware.ResponseTransformerMiddleware())
	// requestLogger := middleware.NewRequestLogger(client)
	// app.Use(requestLogger.LogRequest())

	var (
		httpAddr = ":3000"
	)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		if err := app.Listen(httpAddr); err != nil {
			errs <- err
		}
		internal.StartApp(db, app)
	}()

	fmt.Printf("Go Server is running on port %s", httpAddr)
	fmt.Println("exit", <-errs)
}
