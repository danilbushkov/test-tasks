package news

import "github.com/gofiber/fiber/v2"

func (nh *NewsHandlers) List(c *fiber.Ctx) error {
	return c.SendString("List")
}
