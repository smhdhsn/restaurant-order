package contract

import (
	"github.com/smhdhsn/restaurant-order/internal/model"
)

// OrderRepository is the interface representing order repository or it's mock.
type OrderRepository interface {
	Store(*model.OrderDTO) error
}
