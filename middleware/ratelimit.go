package middleware

import (
	"os"
	"strconv"
	"time"

	"github.com/aishwary11/fiber-tut/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimit() fiber.Handler {
	limit, err := strconv.Atoi(os.Getenv("RATE_LIMIT"))
	if err != nil {
		limit = 5
	}
	return limiter.New(limiter.Config{
		Max:        limit,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return utils.ResponseHelper(c, fiber.StatusTooManyRequests, "Rate limit exceeded", nil)
		},
	})
}
