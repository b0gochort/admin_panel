package service

import (
	"github.com/b0gochort/admin_panel/internal/api_db"
	"github.com/b0gochort/admin_panel/internal/model"
)

type AdminService interface {
	GetCompetent() ([]model.CompetentRes, error)
	ChangeCompetent(elem model.CompetentRes) (model.CompetentRes, error)
	AddCompetent(userId int64) error
}

type Service struct {
	AdminService
}

func NewService(ApiDB *api_db.ApiDB) *Service {
	return &Service{
		AdminService: NewAdminService(ApiDB.AdminAPI),
	}
}
