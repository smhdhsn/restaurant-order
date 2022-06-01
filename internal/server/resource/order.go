package resource

import (
	ospb "github.com/smhdhsn/restaurant-order/internal/protos/order/submission"
)

// OrderResource holds order resource's handlers.
type OrderResource struct {
	OrderSubmit ospb.OrderSubmissionServiceServer
}

// NewOrderResource creates a new user resource with all handlers within itself.
func NewOrderResource(osh ospb.OrderSubmissionServiceServer) *OrderResource {
	return &OrderResource{
		OrderSubmit: osh,
	}
}
