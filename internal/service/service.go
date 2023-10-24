package service

import (
	"github.com/dkshi/banklocsrv"
	"github.com/dkshi/banklocsrv/internal/repository"
)

type Authorization interface {
	CreateUser(user banklocsrv.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Atms interface {
	GetAtms(*map[string]interface{}) (*banklocsrv.AtmsData, error)
}

type Offices interface {
	GetOffices(filter *map[string]interface{}) (*banklocsrv.OfficesData, error)
}

type Service struct {
	Authorization
	Atms
	Offices
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Atms:          NewAtmsService(repo),
		Offices:       NewOfficesService(repo),
	}
}
