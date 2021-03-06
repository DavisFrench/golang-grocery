package db

import (
	"errors"
	"fmt"
	"strings"

	gg "DavisFrench/golang-grocery"
)

// for better output when failing to implement the interface
var _ gg.GroceryService = &GroceryService{}

type GroceryService struct {
	inventory []gg.Produce
}

func NewGroceryService() *GroceryService {

	return &GroceryService{
		// initiate the grocery service with the appropriate inventory
		inventory: gg.InitialInventory,
	}
}

// return index of Produce based by matching on produceId
// returns -1 if not found
func (gs *GroceryService) getIndexFromProduceCode(produceCode string) int {

	for i, produce := range gs.inventory {
		// case insensitivity
		if strings.ToLower(produce.ProduceCode) == strings.ToLower(produceCode) {
			return i
		}
	}

	return -1
}

func (gs *GroceryService) AddProduce(produce gg.Produce) error {

	// add produce if Produce code is not already found to be in inventory
	if index := gs.getIndexFromProduceCode(produce.ProduceCode); index == -1 {

		// decimal place check

		gs.inventory = append(gs.inventory, produce)
		return nil

	} else {
		errMsg := fmt.Sprintf("Procuce code: %s, already in the inventory", produce.ProduceCode)
		return errors.New(errMsg)
	}
}

func (gs *GroceryService) DeleteProduce(produceCode string) error {

	index := gs.getIndexFromProduceCode(produceCode)
	if index != -1 {
		if index == len(gs.inventory)-1 {
			gs.inventory = gs.inventory[:index]
		} else {
			gs.inventory = append(gs.inventory[:index], gs.inventory[index+1:]...)
		}
	}

	return nil
}

func (gs *GroceryService) GetProduceByCode(produceCode string) (*gg.Produce, error) {

	index := gs.getIndexFromProduceCode(produceCode)
	if index != -1 {
		return &gs.inventory[index], nil
	}

	return nil, nil
}

// the error here is simply in the event that an actual db is ever implemented
func (gs *GroceryService) GetAllProduce() ([]gg.Produce, error) {
	return gs.inventory, nil
}
