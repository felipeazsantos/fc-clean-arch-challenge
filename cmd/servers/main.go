package main

import "github.com/felipeazsantos/pos-goexpert/fc-clean-arch-challenge/configs"

func main() {
	conf, err := configs.LoadConfig(".")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	
}