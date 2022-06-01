package contract

import (
	"github.com/smhdhsn/restaurant-order/internal/model"
)

// OrderSubmitService is the interface that order service must implement.
type OrderSubmitService interface {
	Submit(*model.OrderDTO) error
}
