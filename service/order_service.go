package service

import (
	"github.com/kiran/NawabiHouse/model"
	"github.com/kiran/NawabiHouse/repository"
)

func PlaceOrder(order model.Order) error {
	if order.Briyani == "" {
		order.Briyani = "no briyani"
	}
	return repository.CreateOrder(order)
}
