package api_db

import (
	"github.com/b0gochort/admin_panel/internal/model"
	"github.com/restream/reindexer/v3"
)

type AdminAPI interface {
	TakeCompetent() ([]model.CompetentItem, error)
	ChangeCompetent(item model.CompetentItem) (model.CompetentItem, error)
	AddCompetent(userId int64) error
}

type ApiDB struct {
	AdminAPI
}

func NewAPIDB(db *reindexer.Reindexer) *ApiDB {
	return &ApiDB{
		AdminAPI: NewChatApi(db),
	}
}
