package service

import (
	"github.com/pkg/errors"
	"github.com/smhdhsn/restaurant-order/internal/model"

	repositoryContract "github.com/smhdhsn/restaurant-order/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-order/internal/service/contract"
)

// OrderSubmitServ contains repositories that will be used within this service.
type OrderSubmitServ struct {
	iRepo repositoryContract.EdibleInventoryRepository
	oRepo repositoryContract.OrderRepository
}

// NewOrderSubmitService creates a service related to order submittion with it's dependencies.
func NewOrderSubmitService(ir repositoryContract.EdibleInventoryRepository, or repositoryContract.OrderRepository) serviceContract.OrderSubmitService {
	return &OrderSubmitServ{
		iRepo: ir,
		oRepo: or,
	}
}

// Submit is responsible for submiting an order which includes storing order details inside database, and decreasing related component's stocks.
func (s *OrderSubmitServ) Submit(o *model.OrderDTO) (err error) {
	err = s.iRepo.Use(o.FoodID)
	if err != nil {
		return errors.Wrap(err, "error on calling inventory gRPC server")
	}

	err = s.oRepo.Store(o)
	if err != nil {
		return errors.Wrap(err, "error on storing order information inside database")
	}

	return
}
