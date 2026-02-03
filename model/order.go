package model

type CustomerDetails struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type Order struct {
	CustomerDetails CustomerDetails `json:"customer_details"`
	Briyani         string          `json:"briyani"`
	Appetizer       string          `json:"appetizer"`
	Breads          string          `json:"breads"`
}
