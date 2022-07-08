package contract

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-order/internal/repository/entity"
)

// This section holds errors that might happen within repository layer.
var (
	ErrLackOfComponents = errors.New("lack_of_components")
)

// InventoryRepository is the interface representing edible inventory repository or it's mock.
type InventoryRepository interface {
	Use(*entity.Food) error
}
