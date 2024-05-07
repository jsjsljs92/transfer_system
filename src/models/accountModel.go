package models

import "time"

type Account struct {
	ID        uint
	AccID     int
	Balance   float32
	Version   int
	Timestamp time.Time
}
