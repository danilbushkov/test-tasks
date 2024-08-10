package news

import (
	"fmt"

	"github.com/danilbushkov/test-tasks/internal/app/db/models"
	"github.com/danilbushkov/test-tasks/internal/app/errors"
	"github.com/danilbushkov/test-tasks/internal/app/structs"
	"gopkg.in/reform.v1"
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
	tx, err := s.db.Orm().Begin()
	if err != nil {
		return 0, fmt.Errorf("%w: %w", errors.ErrDatabase, err)
	}

	defer tx.Rollback()

	if err = tx.Save(n); err != nil {
		return 0, fmt.Errorf("%w: %w", errors.ErrDatabase, err)
	}

	if news.Categories != nil && len(*news.Categories) != 0 {
		categories := []reform.Struct{}
		for i := 0; i < len(*news.Categories); i++ {
			category := models.NewsCategories{
				NewID:      n.ID,
				CategoryId: (*news.Categories)[i],
			}
			categories = append(categories, &category)
		}
		err = tx.InsertMulti(categories...)
		if err != nil {
			return 0, fmt.Errorf("%w: %w", errors.ErrDatabase, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("%w: %w", errors.ErrDatabase, err)
	}

	return n.ID, nil
}
