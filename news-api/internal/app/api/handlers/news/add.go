package news

import (
	"github.com/danilbushkov/test-tasks/internal/app/models"
	"github.com/gofiber/fiber/v2"
)

func (nh *NewsHandlers) Add(c *fiber.Ctx) error {
	n := &models.News{
		Title:   "test",
		Content: "test",
	}
	if err := nh.appContext.DB.Orm().Save(n); err != nil {
		nh.appContext.Log.Error(err)
	}

	return c.SendString("Add")
}
