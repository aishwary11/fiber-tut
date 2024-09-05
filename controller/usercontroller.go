package controller

import (
	"context"
	"time"

	"github.com/aishwary11/fiber-tut/types"
	"github.com/aishwary11/fiber-tut/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func SignIn(c *fiber.Ctx) error {
	var req types.User
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseHelper(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}
	collection := utils.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user bson.M
	err := collection.FindOne(ctx, bson.M{"email": req.Email, "name": req.Name}).Decode(&user)
	if err != nil {
		return utils.ResponseHelper(c, fiber.StatusUnauthorized, "User not found", nil)
	}
	if !utils.VerifyOTP(req.Otp, user["secret"].(string)) {
		return utils.ResponseHelper(c, fiber.StatusUnauthorized, "Invalid OTP", nil)
	}
	userStruct := types.User{
		ID:    user["_id"].(primitive.ObjectID).Hex(),
		Name:  req.Name,
		Email: req.Email,
	}
	token, err := utils.GenerateToken(userStruct)
	if err != nil {
		return utils.ResponseHelper(c, fiber.StatusInternalServerError, "Failed to generate token", nil)
	}
	return utils.ResponseHelper(c, fiber.StatusOK, "Sign In successful", map[string]string{"token": token})
}

func SignUp(c *fiber.Ctx) error {
	return utils.ResponseHelper(c, fiber.StatusCreated, "Sign Up", nil)
}
