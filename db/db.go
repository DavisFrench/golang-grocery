package db

import (
	"errors"
	"regexp"
	"strings"

	gg "DavisFrench/golang-grocery"
)

// for better output when failing to implement the interface
var _ gg.GroceryService = &GroceryService{}

const (
	PRODUCECODE_REGEX        = `^[a-zA-Z\d]{4}-[a-zA-Z\d]{4}-[a-zA-Z\d]{4}-[a-zA-Z\d]{4}$`
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

	// validate the format of the produce
	valid, err := verifyProduceCodeFormat(produce.ProduceCode)
	if err != nil {
		return err
	}

	if !valid {
		return errors.New(PRODUCECODE_FORMAT_ERROR)
	}

	if index := gs.getIndexFromProduceCode(produce.ProduceCode); index == -1 {

		// decimal place check

		gs.inventory = append(gs.inventory, produce)

	}

	// return error if already inserted (use an else)
	return nil
}

func (gs *GroceryService) DeleteProduce(produceCode string) error {

	// validate produceCode format
	valid, err := verifyProduceCodeFormat(produceCode)
	if err != nil {
		return err
	}

	if !valid {
		return errors.New(PRODUCECODE_FORMAT_ERROR)
	}

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

	// validate produceCode format
	valid, err := verifyProduceCodeFormat(produceCode)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, errors.New(PRODUCECODE_FORMAT_ERROR)
	}

	index := gs.getIndexFromProduceCode(produceCode)
	if index != -1 {
		return &gs.inventory[index], nil
	}

	return nil, nil
}

func (gs *GroceryService) GetAllProduce() ([]gg.Produce, error) {
	return gs.inventory, nil
}
