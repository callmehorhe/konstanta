package service

import (
	"time"

	"github.com/callmehorhe/konstanta/models"
)

func (s *Service) CreateTransaction(transaction models.Transaction) error {
	transaction.Create_time = time.Now().Format("2006-01-02 15:04:05")
	transaction.Update_time = time.Now().Format("2006-01-02 15:04:05")
	transaction.Status = models.StatusNew
	if transaction.User_id < 0 || transaction.Amount < 100 {
		transaction.Status = models.StatusError
	}

	return s.repo.Create(transaction)
}

func (s *Service) GetAllTransactionsByID(id int) ([]models.Transaction, error) {
	return s.repo.GetAllTransactionsByID(id)
}

func (s *Service) GetAllTransactionsByEmail(email string) ([]models.Transaction, error) {
	return s.repo.GetAllTransactionsByEmail(email)
}
