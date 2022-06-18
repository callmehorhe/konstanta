package service

import (
	"errors"
	"time"

	"github.com/callmehorhe/konstanta/models"
)

func (s *Service) CreateTransaction(transaction models.Transaction) error {
	transaction.Create_time = time.Now().Format("2006-01-02 15:04:05")
	transaction.Update_time = time.Now().Format("2006-01-02 15:04:05")
	transaction.Status = models.StatusNew
	if transaction.Amount < 100 {
		transaction.Status = models.StatusError
	}

	return s.repo.Create(transaction)
}

func (s *Service) GetAllTransactionsByID(id int) ([]models.Transaction, error) {
	t, err := s.repo.GetAllTransactionsByID(id)
	if err != nil {
		return nil, err
	}
	if len(t) < 1 {
		return nil, errors.New("no transactions")
	}
	return t, nil
}

func (s *Service) GetAllTransactionsByEmail(email string) ([]models.Transaction, error) {
	t, err := s.repo.GetAllTransactionsByEmail(email)
	if err != nil {
		return nil, err
	}
	if len(t) < 1 {
		return nil, errors.New("no transactions")
	}
	return t, nil
}
