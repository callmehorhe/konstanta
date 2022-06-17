package repository

func (r *Repository) UpdateStatus(id int, status string) error {
	return r.db.Table("transactions").Where("t_id=?", id).Update("status", status).Error
}

func (r *Repository) GetTransactionStatusByID(id int) (string, error) {
	var status string
	err := r.db.Select("status").Table("transactions").Where("t_id=?", id).Take(&status).Error
	return status, err
}
