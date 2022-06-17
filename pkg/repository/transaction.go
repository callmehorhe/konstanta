package repository

import (
	"github.com/callmehorhe/konstanta/models"
)

func (r *Repository) Create(transaction models.Transaction) error {
	return r.db.Table("transactions").Create(&transaction).Error
}

func (r *Repository) GetAllTransactionsByID(id int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Table("transactions").Where("user_id=?", id).Find(&transactions).Error
	return transactions, err
}

func (r *Repository) GetAllTransactionsByEmail(email string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Table("transactions").Where("email=?", email).Find(&transactions).Error
	return transactions, err
}
