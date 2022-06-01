package model

import (
	"time"
)

// Order represents the order table's model.
type Order struct {
	ID        uint32 `gorm:"primaryKey"`
	UserID    uint32 `gorm:"index"`
	FoodID    uint32 `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// OrderDTO represents order's data transfer object.
type OrderDTO struct {
	ID        uint32
	UserID    uint32
	FoodID    uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}

// OrderListDTO represents a list of OrderDTOs.
type OrderListDTO []*OrderDTO
