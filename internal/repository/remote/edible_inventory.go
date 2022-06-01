package remote

import (
	"context"

	"github.com/pkg/errors"

	eipb "github.com/smhdhsn/restaurant-order/internal/protos/edible/inventory"
	repositoryContract "github.com/smhdhsn/restaurant-order/internal/repository/contract"
)

// EdibleInventoryRepo contains repository's database connection.
type EdibleInventoryRepo struct {
	eic eipb.EdibleInventoryServiceClient
	ctx *context.Context
}

// NewEdibleInventoryRepository creates an instance of the remote repository with gRPC connection.
func NewEdibleInventoryRepository(ctx *context.Context, conn eipb.EdibleInventoryServiceClient) repositoryContract.EdibleInventoryRepository {
	return &EdibleInventoryRepo{
		eic: conn,
		ctx: ctx,
	}
}

// Use is responsible for decreasing stocks of a component.
func (s *EdibleInventoryRepo) Use(foodID uint32) error {
	req := eipb.InventoryUseRequest{
		FoodId: foodID,
	}

	_, err := s.eic.Use(*s.ctx, &req)
	if err != nil {
		return errors.Wrap(err, "error on calling use on edible gRPC server")
	}

	return nil
}
