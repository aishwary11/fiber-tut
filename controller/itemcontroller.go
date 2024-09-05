package controller

import (
	"strconv"

	"github.com/aishwary11/fiber-tut/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAllItems(c *fiber.Ctx) error {
	return utils.ResponseHelper(c, fiber.StatusOK, "Get All Items", utils.Items)
}

func GetItemByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for _, item := range utils.Items {
		if item.ID == id {
			return utils.ResponseHelper(c, fiber.StatusOK, "Item found", item)
		}
	}
	return utils.ResponseHelper(c, fiber.StatusNotFound, "Item not found", nil)
}
