package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"

	"fiber-api-example/app/config"
)

func Create() *fiber.App {
	//database.SetupDatabase()

	app := fiber.New(config.GetFiberConfig())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	//setupMiddlewares(app)

	return app
}

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint                            // wait for OS signal

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := a.Listen(viper.GetString("APP_ADDR")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Run server.
	if err := a.Listen(viper.GetString("APP_ADDR")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
