package data

import "time"

type Transaction struct {
	Amount    float32   `json:"amount"`
	Date      time.Time `json:"time"`
	Necessity int8      `json:"necessity"`
	Category  string    `json:"category"`
	Reason    string    `json:"reason"`
}

