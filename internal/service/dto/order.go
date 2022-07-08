package dto

import (
	"time"
)

// Order represents order's data transfer object.
type Order struct {
	ID        uint32
	UserID    uint32
	FoodID    uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}
