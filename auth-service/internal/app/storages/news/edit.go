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
	tx, err := s.db.Orm().Begin()
	if err != nil {
		return fmt.Errorf("%w: %w", e.ErrDatabase, err)
	}

	defer tx.Rollback()

	if len(columns) != 0 {
		if err := tx.UpdateColumns(n, columns...); err != nil {
			if err == reform.ErrNoRows {
				return errors.New("News does not exist")
			} else {
				return fmt.Errorf("%w: %w", e.ErrDatabase, err)
			}
		}
	}
	if news.Categories != nil {
		placeholder := s.db.Orm().Placeholder(1)
		tail := fmt.Sprintf("WHERE news_id = %s", placeholder)
		if _, err = tx.DeleteFrom(
			models.NewsCategoriesView,
			tail,
			n.ID,
		); err != nil {
			return fmt.Errorf("%w: %w", e.ErrDatabase, err)
		}
		if len(*news.Categories) != 0 {
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
				return fmt.Errorf("%w: %w", e.ErrDatabase, err)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%w: %w", e.ErrDatabase, err)
	}

	return nil
}
