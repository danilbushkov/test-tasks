package news

import (
	"errors"
	"fmt"

	"github.com/danilbushkov/test-tasks/internal/app/db/models"
	e "github.com/danilbushkov/test-tasks/internal/app/errors"
	"github.com/danilbushkov/test-tasks/internal/app/structs"
	"gopkg.in/reform.v1"
)

func (s *NewsStorage) Edit(news *structs.News) error {
	columns := []string{}
	if err := news.CheckId(); err != nil {
		return err
	}
	n := &models.News{ID: *news.Id}
	if news.Title != nil {
		if err := news.CheckTitle(); err != nil {
			return err
		}
		n.Title = *news.Title
		columns = append(columns, "title")
	}
	if news.Content != nil {
		if err := news.CheckContent(); err != nil {
			return err
		}
		n.Content = *news.Content
		columns = append(columns, "content")
	}

	if err := s.db.Orm().UpdateColumns(n, columns...); err != nil {
		if err == reform.ErrNoRows {
			return errors.New("News does not exist")
		} else {
			return fmt.Errorf("%w: %w", e.ErrDatabase, err)
		}
	}

	return nil
}
