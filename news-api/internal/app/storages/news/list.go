package news

import (
	"fmt"

	"github.com/danilbushkov/test-tasks/internal/app/db/models"
	"github.com/danilbushkov/test-tasks/internal/app/errors"
	"github.com/danilbushkov/test-tasks/internal/app/structs"
	"gopkg.in/reform.v1"
)

func (s *NewsStorage) List() ([]*structs.News, error) {
	rows, err := s.db.Orm().SelectRows(models.NewsTable, "")
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrDatabase, err)
	}
	defer rows.Close()

	list := []*structs.News{}
	for {
		var news models.News
		if err = s.db.Orm().NextRow(&news, rows); err != nil {
			break
		}
		categories := []int64{}
		r, err := s.db.Orm().FindRows(
			models.NewsCategoriesView,
			"news_id", news.ID)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", errors.ErrDatabase, err)
		}
		defer r.Close()

		for {
			var category models.NewsCategories
			if err = s.db.Orm().NextRow(&category, r); err != nil {
				break
			}
			categories = append(categories, category.CategoryId)
		}
		if err != reform.ErrNoRows {
			return nil, fmt.Errorf("%w: %w", errors.ErrDatabase, err)
		}
		list = append(list, &structs.News{
			Id:         &news.ID,
			Title:      &news.Title,
			Content:    &news.Content,
			Categories: &categories,
		})
	}
	if err != reform.ErrNoRows {
		return nil, fmt.Errorf("%w: %w", errors.ErrDatabase, err)
	}
	return list, nil
}
