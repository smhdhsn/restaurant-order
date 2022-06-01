package mysql

import (
	"gorm.io/gorm"

	"github.com/pkg/errors"
	"github.com/smhdhsn/restaurant-order/internal/model"

	repositoryContract "github.com/smhdhsn/restaurant-order/internal/repository/contract"
)

// OrderRepo contains repository's database connection.
type OrderRepo struct {
	model model.Order
	db    *gorm.DB
}

// NewOrderRepo creates an instance of the repository with database connection.
func NewOrderRepository(db *gorm.DB, m model.Order) repositoryContract.OrderRepository {
	return &OrderRepo{
		model: m,
		db:    db,
	}
}

// Store is responsible for storing order's details into database.
func (r *OrderRepo) Store(o *model.OrderDTO) error {
	err := r.db.Model(r.model).Create(o).Error
	if err != nil {
		return errors.Wrap(err, "error on storing order inside database")
	}

	return nil
}
