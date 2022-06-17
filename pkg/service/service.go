package service

import "github.com/callmehorhe/konstanta/pkg/repository"


type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}