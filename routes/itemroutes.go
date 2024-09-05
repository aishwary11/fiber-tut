package routes

import (
	"github.com/aishwary11/fiber-tut/controller"
	"github.com/gofiber/fiber/v2"
)

func ItemRoutes(app *fiber.App) {
	ItemGroup := app.Group("/item")
	ItemGroup.Get("/", controller.GetAllItems)
	ItemGroup.Get("/:id", controller.GetItemByID)
}
