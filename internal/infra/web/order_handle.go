package web

import (
	"encoding/json"
	"net/http"

	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/entity"
	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/usecase"
)

type WebOrderHadler struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewWebOrderHandler(OrderRepository entity.OrderRepositoryInterface) *WebOrderHadler {
	return &WebOrderHadler{OrderRepository: OrderRepository}
}

func (h *WebOrderHadler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrderUseCase := usecase.NewCreateOrderUseCase(h.OrderRepository)
	output, err := createOrderUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *WebOrderHadler) ListOrders(w http.ResponseWriter, r *http.Request) {
	listOrdersUseCase := usecase.NewListOrdersUseCase(h.OrderRepository)
	output, err := listOrdersUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
