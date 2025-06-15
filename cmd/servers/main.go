package main

import (
	"fmt"
	"log"

	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/configs"
	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/infra/database"
	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/infra/web"
	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/infra/web/webserver"
	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/usecase"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("failed to load config: " + err.Error())
	}

	db, err := configs.ConnectDB(cfg)
	if err != nil {
		log.Fatal("cannot connect with DB:", err)
	}

	orderRepository := database.NewOrderRepository(db)
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository)
	listOrdersUseCase := usecase.NewListOrdersUseCase(orderRepository)

	webserver := webserver.NewWebServer(cfg.WebServerPort)
	webOrderHandler := web.NewWebOrderHandler(orderRepository)
	webserver.AddHandler("/orders/create", webOrderHandler.Create)
	webserver.AddHandler("/orders/list", webOrderHandler.ListOrders)

	fmt.Println("Starting web server on port:", cfg.WebServerPort)
	go webserver.Start()

}
