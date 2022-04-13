package data

import "time"

type TransactionType uint8

const (
	Income TransactionType = iota
	Expense
	ExpectedIncome
)

type Transaction struct {
	Amount       float32         `json:"amount"`
	Date         time.Time       `json:"time"`
	Necessity    int8            `json:"necessity"`
	Type         TransactionType `json:"transaction_type"`
	MainCategory []string        `json:"main_category"`
	SubCategory  string          `json:"sub_category"`
	Reason       string          `json:"reason"`
}
