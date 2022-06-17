package models

const (
	StatusNew        = "НОВЫЙ"
	StatusSuccess    = "УСПЕХ"
	StatusNotSuccess = "НЕУСПЕХ"
	StatusError      = "ОШИБКА"
	StatusCanceled   = "ОТМЕНЕН"
)

type Transaction struct {
	T_id        int     `json:"t_id"`
	User_id     int     `json:"user_id"`
	Email       string  `json:"email"`
	Amount      float64 `json:"amount"`
	Сurrency    string  `json:"currency"`
	Create_time string  `json:"create_time"`
	Update_time string  `json:"update_time"`
	Status      string  `json:"status"`
}
