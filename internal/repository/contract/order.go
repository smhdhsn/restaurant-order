package contract

import (
	"github.com/smhdhsn/restaurant-order/internal/repository/entity"
)

// OrderRepository is the interface representing order repository or it's mock.
type OrderRepository interface {
	Store(*entity.Order) error
}
