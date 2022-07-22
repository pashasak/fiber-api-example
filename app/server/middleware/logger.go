package middleware

import (
	"fiber-api-example/app/utils/logger"
	"github.com/gofiber/fiber/v2"
	"sync"
)

type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next: nil,
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}
	// Override default config
	cfg := config[0]
	// Set default values
	if cfg.Next == nil {
		cfg.Next = ConfigDefault.Next
	}
	return cfg
}

func Logger(config ...Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config...)

	// Set variables
	var (
		once       sync.Once
		errHandler fiber.ErrorHandler
	)

	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// Set error handler once
		once.Do(func() {
			// override error handler
			errHandler = c.App().ErrorHandler
		})

		// Handle request, store err for logging
		chainErr := c.Next()

		// Manually call error handler
		if chainErr != nil {
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		logger.Info(chainErr.Error())
		logger.Sync()

		return nil
	}
}
