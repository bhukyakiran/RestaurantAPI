package repository

import (
	"fmt"
	"github.com/kiran/NawabiHouse/connection"
	"github.com/kiran/NawabiHouse/model"
)

func CreateOrder(order model.Order) error {
	query := `INSERT INTO orders (name, phone_number, briyani, appetizer, breads) VALUES ($1, $2, $3, $4, $5)`
	_, err := connection.DB.Exec(query,
		order.CustomerDetails.Name,
		order.CustomerDetails.PhoneNumber,
		order.Briyani,
		order.Appetizer,
		order.Breads,
	)
	if err != nil {
		return fmt.Errorf("failed to create order: %v", err)
	}
	return nil
}
