package routes

import (
	"github.com/aishwary11/fiber-tut/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/user")
	userGroup.Post("/signin", controller.SignIn)
	userGroup.Post("/signup", controller.SignUp)
}
