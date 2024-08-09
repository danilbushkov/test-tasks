package news

import "github.com/gofiber/fiber/v2"

func (nh *NewsHandlers) Edit(c *fiber.Ctx) error {
	return c.SendString("Edit")
}
