package entity

import (
	"time"
)

// Order represents the order repository's entity.
type Order struct {
	ID        uint32
	UserID    uint32
	FoodID    uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}
