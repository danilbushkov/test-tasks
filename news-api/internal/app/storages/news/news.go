package news

import (
	"github.com/danilbushkov/test-tasks/internal/app/context"
	"github.com/danilbushkov/test-tasks/internal/app/db"
	"github.com/sirupsen/logrus"
)

type NewsStorage struct {
	db  *db.DB
	log *logrus.Logger
}

func New(db *db.DB, log *logrus.Logger) *NewsStorage {
	return &NewsStorage{
		db,
		log,
	}
}

func NewFromAppContext(actx *context.AppContext) *NewsStorage {
	return &NewsStorage{
		db:  actx.DB,
		log: actx.Log,
	}
}
