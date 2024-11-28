package user

import (
	"math"
	"time"
)

type Balance int32

// FromCents converts the balance from cents to higher currency with 2 decimal places
func (b *Balance) FromCents() float64 {
	return math.Round(float64(*b)/100*100) / 100
}

type User struct {
	ID      uint32
	Balance Balance
}

type Transaction struct {
	ID        string
	UserID    uint32
	Amount    Balance
	CreatedAt time.Time
}
