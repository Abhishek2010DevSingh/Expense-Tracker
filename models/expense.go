package models

import "time"

type Expense struct {
	ID          uint
	Date        time.Time
	Description string
	Amount      float64
}
