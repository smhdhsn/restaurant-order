package contract

import (
	"github.com/smhdhsn/restaurant-order/internal/service/dto"
)

// SubmissionService is the interface that order service must implement.
type SubmissionService interface {
	Submit(*dto.Order) error
}
