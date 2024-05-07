package models

import "time"

type Payout struct {
	ID        uint
	FromAccID int
	Amount    float32
	Timestamp time.Time
}
