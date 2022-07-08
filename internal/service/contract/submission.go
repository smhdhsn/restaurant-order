package contract

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-order/internal/service/dto"
)

// This section holds errors that might happen within service layer.
var (
	ErrLackOfComponents = errors.New("lack_of_components")
)

// SubmissionService is the interface that order service must implement.
type SubmissionService interface {
	Submit(*dto.Order) error
}
