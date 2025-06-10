package main

import (
	"log"

	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/configs"
	"github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/internal/usecase"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	db, err := configs.ConnectDB(cfg)
	if err != nil {
		log.Fatal("cannot connect with DB:", err)
	}

	 _ = usecase.NewListOrdersUseCase()

}
