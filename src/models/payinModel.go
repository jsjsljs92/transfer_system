package models

import "time"

type Payin struct {
	ID        uint
	ToAccID   int
	Amount    float32
	Timestamp time.Time
}
