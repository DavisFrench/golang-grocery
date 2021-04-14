package db

import (
	"errors"
	"regexp"

	gg "DavisFrench/golang-grocery"
)

// for better output when failing to implement the interface
var _ gg.GroceryService = &GroceryService{}

const (
	PRODUCECODE_REGEX = `^[a-zA-Z\d]{4}-[a-zA-Z\d]{4}-[a-zA-Z\d]{4}-[a-zA-Z\d]{4}$`
	PRODUCECODE_FORMAT_ERROR = "The produce_id should be in the following format: xxxx-xxxx-xxxx-xxxx, where x is an alphanumeric character and case insensitive"
)

type GroceryService struct {
	inventory []gg.Produce
}

func NewGroceryService() *GroceryService {

	var inventory []gg.Produce

	return &GroceryService{
		inventory: inventory,
	}
}

func verifyProduceCodeFormat(produceCode string) (bool, error) {
	return regexp.MatchString(PRODUCECODE_REGEX, produceCode)
}

func (gs *GroceryService) AddProduce(produce gg.Produce) error {

	// validate the format of the produce

	gs.inventory = append(gs.inventory, produce)
	return nil
}

func (gs *GroceryService) DeleteProduce(produceCode string) error {

	// validate produceCode format
	valid, err := verifyProduceCodeFormat(produceCode)
	if err != nil{
		return err
	}

	if !valid {
		return errors.New(PRODUCECODE_FORMAT_ERROR)
	}

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

	// TODO: add case insensitivity for any this and the delete

	// validate produceCode format
	valid, err := verifyProduceCodeFormat(produceCode)
	if err != nil{
		return nil, err
	}

	if !valid {
		return nil, errors.New(PRODUCECODE_FORMAT_ERROR)
	}

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
