package resource

import (
	submissionProto "github.com/smhdhsn/restaurant-order/internal/protos/order/submission"
)

// OrderResource holds order resource's handlers.
type OrderResource struct {
	SubmissionHandler submissionProto.OrderSubmissionServiceServer
}

// NewOrderResource creates a new user resource with all handlers within itself.
func NewOrderResource(sh submissionProto.OrderSubmissionServiceServer) *OrderResource {
	return &OrderResource{
		SubmissionHandler: sh,
	}
}
