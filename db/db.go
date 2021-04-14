package db

import (
	gg "DavisFrench/golang-grocery"
)

// for better output when failing to implement the interface
var _ gg.GroceryService = &GroceryService{}

type GroceryService struct {
	inventory []gg.Produce
}

func NewGroceryService() *GroceryService {

	var inventory []gg.Produce

	return &GroceryService{
		inventory: inventory,
	}
}

func (gs *GroceryService) AddProduce(produce gg.Produce) error {

	// validate the format of the produce

	gs.inventory = append(gs.inventory, produce)
	return nil
}

func (gs *GroceryService) DeleteProduce(produceCode string) error {

	// validate produceCode format

	for i, produce := range gs.inventory {
		if produce.ProduceCode == produceCode {
			if i == len(gs.inventory) {
				gs.inventory = gs.inventory[:i]
			} else {
				gs.inventory = append(gs.inventory[:i], gs.inventory[i+1:]...)
			}
		}
	}

	return nil
}

func (gs *GroceryService) GetProduceByCode(produceCode string) (*gg.Produce, error) {

	// validate produceCode format

	for _, produce := range gs.inventory {
		if produce.ProduceCode == produceCode {
			return &produce, nil
		}
	}

	return nil, nil
}

func (gs *GroceryService) GetAllProduce() ([]gg.Produce, error) {
	return gs.inventory, nil
}
