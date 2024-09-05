package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger(port string) fiber.Handler {
	return logger.New(logger.Config{
		Format: fmt.Sprintf("[${time}] [${ip}]: %s ${status} - ${method} ${path} - ${body}\n", port),
		Output: os.Stdout,
	})
}
