package controller

import "github.com/gofiber/fiber/v2"

func GetReport(c *fiber.Ctx) error {
	var res string
	userId := c.Params("userId")

	if userId != "" {
		res = userId
	}

	return c.JSON(res)
}
