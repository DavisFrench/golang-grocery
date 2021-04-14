package db

import (
	gg "DavisFrench/golang-grocery"
)

type GroceryService struct {
	inventory []gg.Produce
}

func NewGroceryService() GroceryService {

	var inventory []gg.Produce

	return GroceryService{
		inventory: inventory,
	}	
}
