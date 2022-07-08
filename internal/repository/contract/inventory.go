package contract

import (
	"github.com/smhdhsn/restaurant-order/internal/repository/entity"
)

// InventoryRepository is the interface representing edible inventory repository or it's mock.
type InventoryRepository interface {
	Use(*entity.Food) error
}
