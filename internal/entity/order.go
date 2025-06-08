package entity

type Order struct {
	ID string
	Price float64
	Tax float64
	FinalPrice float64
}

func NewOrder(id string, price, tax float64) *Order {
	return &Order{
		ID:         id,
		Price:      price,
		Tax:        tax,
		FinalPrice: price + tax,
	}
}

