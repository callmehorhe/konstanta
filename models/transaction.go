package models

const (
	StatusNew        = "НОВЫЙ"
	StatusSuccess    = "УСПЕХ"
	StatusNotSuccess = "НЕУСПЕХ"
	StatusError      = "ОШИБКА"
	StatusCanceled   = "ОТМЕНЕН"
)

type Transaction struct {
	T_id        int     `json:"t_id"        gorm:"primaryKey;column:t_id"`
	User_id     int     `json:"user_id"     gorm:"column:user_id"`
	Email       string  `json:"email"       gorm:"column:email"`
	Amount      float64 `json:"amount"      gorm:"column:amount"`
	Сurrency    string  `json:"currency"    gorm:"column:currency"`
	Create_time string  `json:"create_time" gorm:"column:create_time"`
	Update_time string  `json:"update_time" gorm:"column:update_time"`
	Status      string  `json:"status"      gorm:"column:status"`
}
