package models

import "time"

type Transaction struct {
	ID        int          `json:"id" gorm:"primaryKey:autoIncrement"`
	UserID    int          `json:"user_id"`
	User      UserResponse `json:"user"`
	StartDate time.Time    `json:"start_date"`
	DueDate   time.Time    `json:"due_date"`
	Price     int          `json:"price"`
	Status    string       `json:"status"`
}

type TransactionResponse struct {
	ID        int          `json:"id"`
	UserID    int          `json:"user_id"`
	User      UserResponse `json:"user"`
	StartDate time.Time    `json:"start_date"`
	DueDate   time.Time    `json:"due_date"`
	Price     int          `json:"price"`
	Status    string       `json:"status"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
