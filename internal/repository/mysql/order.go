package mysql

import (
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-order/internal/repository/entity"

	repositoryContract "github.com/smhdhsn/restaurant-order/internal/repository/contract"
)

// order represents the order table's model.
type order struct {
	ID        uint32 `gorm:"primaryKey"`
	UserID    uint32 `gorm:"index"`
	FoodID    uint32 `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// OrderRepo contains repository's database connection.
type OrderRepo struct {
	model order
	db    *gorm.DB
}

// NewOrderRepository creates an instance of the repository with database connection.
func NewOrderRepository(db *gorm.DB) repositoryContract.OrderRepository {
	return &OrderRepo{
		model: order{},
		db:    db,
	}
}

// Store is responsible for storing order's details into database.
func (r *OrderRepo) Store(oEntity *entity.Order) error {
	oModel := singleOrderEntityToModel(oEntity)

	err := r.db.Model(r.model).Create(oModel).Error
	if err != nil {
		return errors.Wrap(err, "error on storing order inside database")
	}

	return nil
}

// singleOrderEntityToModel is responsible for transforming an order enity into order model struct.
func singleOrderEntityToModel(oEntity *entity.Order) *order {
	return &order{
		ID:        oEntity.ID,
		UserID:    oEntity.UserID,
		FoodID:    oEntity.FoodID,
		CreatedAt: oEntity.CreatedAt,
		UpdatedAt: oEntity.UpdatedAt,
	}
}
