package news

import (
	"errors"
	"fmt"
	"github.com/danilbushkov/test-tasks/internal/app/db/models"
	e "github.com/danilbushkov/test-tasks/internal/app/errors"
	"gopkg.in/reform.v1"
)

func (s *NewsStorage) Delete(id int64) error {

	n := &models.News{
		ID: id,
	}
	if err := s.db.Orm().Delete(n); err != nil {
		if err == reform.ErrNoRows {
			return errors.New("News does not exist")
		} else {
			return fmt.Errorf("%w: %w", e.ErrDatabase, err)
		}
	}

	return nil
}
