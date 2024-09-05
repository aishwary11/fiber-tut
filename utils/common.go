package utils

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/aishwary11/fiber-tut/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp/totp"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(user types.User) (string, error) {
	claims := jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ResponseHelper(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := fiber.Map{
		"message": message,
		"data":    data,
		"status":  statusCode < fiber.StatusBadRequest,
	}
	return c.Status(statusCode).JSON(response)
}

func GenerateTOTPSecret() string {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer: "Aishwary TOTP",
	})
	if err != nil {
		log.Fatalf("Failed to generate TOTP key: %v", err)
	}
	secret := key.Secret()
	if len(secret) > 24 {
		secret = secret[:24]
	} else if len(secret) < 24 {
		secret = secret + strings.Repeat("A", 24-len(secret))
	}
	return secret
}

func VerifyOTP(otp string, secret string) bool {
	return totp.Validate(otp, secret)
}
