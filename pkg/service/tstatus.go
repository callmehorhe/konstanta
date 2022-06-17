package service

import (
	"errors"

	"github.com/callmehorhe/konstanta/models"
)

func (s *Service) UpdateStatus(id int, status string) error {
	oldStatus, err := s.repo.GetTransactionStatusByID(id)
	if err != nil {
		return err
	}

	if oldStatus == models.StatusSuccess || oldStatus == models.StatusNotSuccess {
		return errors.New("can't change status")
	}

	return s.repo.UpdateStatus(id, status)
}

func (s *Service) GetStatusByID(id int) (string, error) {
	return s.repo.GetTransactionStatusByID(id)
}