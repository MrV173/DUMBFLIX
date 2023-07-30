package transactiondto

import "time"

type CreateTransaction struct {
	UserID    int       `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	DueDate   time.Time `json:"due_date"`
	Price     int       `json:"price"`
	Status    string    `json:"status"`
}
