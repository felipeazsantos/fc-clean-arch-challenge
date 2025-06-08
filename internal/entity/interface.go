package entity

type OrderRepositoryInterface interface {
	ListOrders() ([]Order, error)
}