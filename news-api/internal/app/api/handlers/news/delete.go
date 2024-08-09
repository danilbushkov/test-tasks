package news

import "github.com/gofiber/fiber/v2"

func (nh *NewsHandlers) Delete(c *fiber.Ctx) error {
	return c.SendString("Edit")
}
