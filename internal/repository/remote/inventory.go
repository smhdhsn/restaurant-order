package remote

import (
	"context"

	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-order/internal/repository/entity"

	inventoryProto "github.com/smhdhsn/restaurant-order/internal/protos/edible/inventory"
	repositoryContract "github.com/smhdhsn/restaurant-order/internal/repository/contract"
)

// InventoryRepo contains repository's database connection.
type InventoryRepo struct {
	client inventoryProto.EdibleInventoryServiceClient
	ctx    *context.Context
}

// NewInventoryRepository creates an instance of the remote repository with gRPC connection.
func NewInventoryRepository(ctx *context.Context, conn inventoryProto.EdibleInventoryServiceClient) repositoryContract.InventoryRepository {
	return &InventoryRepo{
		client: conn,
		ctx:    ctx,
	}
}

// Use is responsible for decreasing stocks of a component.
func (s *InventoryRepo) Use(fEntity *entity.Food) error {
	req := singleFoodEntityToProtoUseReq(fEntity)

	_, err := s.client.Use(*s.ctx, req)
	if err != nil {
		return errors.Wrap(err, "error on calling use on edible gRPC server")
	}

	return nil
}

// singleFoodEntityToProtoUseReq is responsible for transforming a food entity into use request struct.
func singleFoodEntityToProtoUseReq(fEntity *entity.Food) *inventoryProto.InventoryUseRequest {
	return &inventoryProto.InventoryUseRequest{
		FoodId: fEntity.FoodID,
	}
}
