package service

import (
	"github.com/dkshi/banklocsrv"
	"github.com/dkshi/banklocsrv/internal/repository"
)

type AtmsService struct {
	repo repository.Atms
}

func NewAtmsService(repo repository.Atms) *AtmsService {
	return &AtmsService{repo: repo}
}

func (s *AtmsService) GetAtms(filter *map[string]interface{}) (*banklocsrv.AtmsData, error) {
	atms, err := s.repo.GetAtms(filter)

	if err != nil {
		return &banklocsrv.AtmsData{}, err
	}

	return atms, nil
}

