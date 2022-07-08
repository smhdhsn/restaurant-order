package service

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-order/internal/repository/entity"
	"github.com/smhdhsn/restaurant-order/internal/service/dto"

	repositoryContract "github.com/smhdhsn/restaurant-order/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-order/internal/service/contract"
)

// SubmissionServ contains repositories that will be used within this service.
type SubmissionServ struct {
	iRepo repositoryContract.InventoryRepository
	oRepo repositoryContract.OrderRepository
}

// NewSubmissionService creates a service related to order submission with it's dependencies.
func NewSubmissionService(
	ir repositoryContract.InventoryRepository,
	or repositoryContract.OrderRepository,
) serviceContract.SubmissionService {
	return &SubmissionServ{
		iRepo: ir,
		oRepo: or,
	}
}

// Submit is responsible for submiting an order which includes storing order details inside database, and decreasing related component's stocks.
func (s *SubmissionServ) Submit(oDTO *dto.Order) (err error) {
	fEntity := singleOrderDTOToFoodEntity(oDTO)

	err = s.iRepo.Use(fEntity)
	if err != nil {
		if errors.Is(err, repositoryContract.ErrLackOfComponents) {
			return serviceContract.ErrLackOfComponents
		}

		return errors.Wrap(err, "error on calling inventory gRPC server")
	}

	oEntity := singleOrderDTOToEntity(oDTO)

	err = s.oRepo.Store(oEntity)
	if err != nil {
		return errors.Wrap(err, "error on storing order information inside database")
	}

	return
}

// singleOrderDTOToFoodEntity is responsible for transforming an order dto into food entity struct.
func singleOrderDTOToFoodEntity(oDTO *dto.Order) *entity.Food {
	return &entity.Food{
		FoodID: oDTO.FoodID,
	}
}

// singleOrderDTOToEntity is responsible for transforming an order dto into order entity struct.
func singleOrderDTOToEntity(oDTO *dto.Order) *entity.Order {
	return &entity.Order{
		ID:        oDTO.ID,
		UserID:    oDTO.UserID,
		FoodID:    oDTO.FoodID,
		CreatedAt: oDTO.CreatedAt,
		UpdatedAt: oDTO.UpdatedAt,
	}
}
