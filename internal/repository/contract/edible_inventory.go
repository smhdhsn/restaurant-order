package contract

// EdibleInventoryRepository is the interface representing edible inventory repository or it's mock.
type EdibleInventoryRepository interface {
	Use(foodID uint32) error
}
