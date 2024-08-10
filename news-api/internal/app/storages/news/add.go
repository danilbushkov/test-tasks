package news

import (
	"fmt"

	"github.com/danilbushkov/test-tasks/internal/app/db/models"
	"github.com/danilbushkov/test-tasks/internal/app/errors"
	"github.com/danilbushkov/test-tasks/internal/app/structs"
)

func (s *NewsStorage) Add(news *structs.News) (int64, error) {
	if err := news.CheckTitle(); err != nil {
		return 0, err
	}
	if err := news.CheckContent(); err != nil {
		return 0, err
	}
	n := &models.News{
		Title:   *news.Title,
		Content: *news.Content,
	}
	if err := s.db.Orm().Save(n); err != nil {
		return 0, fmt.Errorf("%w: %w", errors.ErrDatabase, err)
	}

	return n.ID, nil
}
