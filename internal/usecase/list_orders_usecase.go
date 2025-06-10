package usecase

import "github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (u *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := u.OrderRepository.ListOrders()
	if err != nil {
		return nil, err
	}

	var result []OrderOutputDTO
	for _, order := range orders {
		result = append(result, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return result, nil
}