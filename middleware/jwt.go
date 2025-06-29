package middleware

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/aishwary11/fiber-tut/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return utils.ResponseHelper(c, fiber.StatusUnauthorized, "Missing or malformed JWT", nil)
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			return utils.ResponseHelper(c, fiber.StatusUnauthorized, "Invalid or expired JWT", nil)
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return utils.ResponseHelper(c, fiber.StatusUnauthorized, "Invalid token claims", nil)
		}
		email, emailOk := claims["email"].(string)
		name, nameOk := claims["name"].(string)
		if !emailOk || !nameOk {
			return utils.ResponseHelper(c, fiber.StatusUnauthorized, "Invalid token claims", nil)
		}
		collection := utils.GetCollection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var user bson.M
		if err := collection.FindOne(ctx, bson.M{"email": email, "name": name}).Decode(&user); err != nil {
			return utils.ResponseHelper(c, fiber.StatusUnauthorized, "User not found", nil)
		}
		return c.Next()
	}
}
