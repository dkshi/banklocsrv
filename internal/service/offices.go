package service

import (
	"github.com/dkshi/banklocsrv"
	"github.com/dkshi/banklocsrv/internal/repository"
)

type OfficesService struct {
	repo repository.Offices
}

func NewOfficesService(repo repository.Offices) *OfficesService {
	return &OfficesService{repo: repo}
}

func (s *OfficesService) GetOffices(filter *map[string]interface{}) (*banklocsrv.OfficesData, error) {
	offices, err := s.repo.GetOffices(filter)

	if err != nil {
		return &banklocsrv.OfficesData{}, err
	}

	return offices, nil
}
