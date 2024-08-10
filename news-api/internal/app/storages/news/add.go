package news

import (
	"fmt"

	"github.com/danilbushkov/test-tasks/internal/app/db/models"
	"github.com/danilbushkov/test-tasks/internal/app/errors"
	"github.com/danilbushkov/test-tasks/internal/app/structs"
)

func (s *NewsStorage) Add(news *structs.News) error {
	if err := news.CheckTitle(); err != nil {
		return err
	}
	if err := news.CheckContent(); err != nil {
		return err
	}
	n := &models.News{
		Title:   *news.Title,
		Content: *news.Content,
	}
	if err := s.db.Orm().Save(n); err != nil {
		return fmt.Errorf("%w: %w", errors.ErrDatabase, err)
	}

	return nil
}
