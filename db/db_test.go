package db

import (
	"testing"

	gg "DavisFrench/golang-grocery"
)

// API checks that the produce_code is in the proper format
func Test_GetProduceByProduceCode(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		errExpected    bool
		returnExpected bool
	}{
		{
			name:           "Produce not found",
			input:          "asdf-asdf-asdf-asdf",
			errExpected:    false,
			returnExpected: false,
		},
		{
			name:           "Produce found",
			input:          "TQ4C-VV6T-75ZX-1RMR",
			errExpected:    false,
			returnExpected: true,
		},
	}

	for _, test := range tests {

		// can insert a custom inventory if needed
		GS := GroceryService{
			inventory: gg.InitialInventory,
		}

		t.Run(test.name, func(t *testing.T) {

			produce, err := GS.GetProduceByCode(test.input)

			if (err != nil) != test.errExpected {
				t.Errorf("errExpected: %v, errReceived: %v", test.errExpected, (err != nil))
			}

			if (produce != nil) != test.returnExpected {
				t.Errorf("returnExpected: %v, returnReceived: %v", test.returnExpected, (produce != nil))
			}
		})
	}
}

func Test_DeleteProduce(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		errExpected bool
		lenExpected int
	}{
		{
			name:        "Successful delete",
			input:       "TQ4C-VV6T-75ZX-1RMR",
			errExpected: false,
			lenExpected: 3,
		},
	}

	for _, test := range tests {

		// can insert a custom inventory if needed
		GS := GroceryService{
			inventory: gg.InitialInventory,
		}

		t.Run(test.name, func(t *testing.T) {

			err := GS.DeleteProduce(test.input)

			if (err != nil) != test.errExpected {
				t.Errorf("errExpected: %v, errReceived: %v", test.errExpected, (err != nil))
			}

			if test.lenExpected != len(GS.inventory) {
				t.Errorf("lenExpected: %d, lenFound: %d", test.lenExpected, len(GS.inventory))
			}
		})
	}
}

/*
func Test_AddProduce(t *testing.T) {
	tests := []struct {
		name string
		errExpected bool
		returnExpected bool
	}{
		{
			name: "",
			errExpected: true,
			returnExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}*/

func Test_GetAllProduce(t *testing.T) {
	tests := []struct {
		name           string
		errExpected    bool
		returnExpected bool
		lenExpected    int
	}{
		{
			name:           "Successful return",
			errExpected:    false,
			returnExpected: true,
			lenExpected:    4,
		},
	}

	for _, test := range tests {

		// can insert a custom inventory if needed
		GS := GroceryService{
			inventory: gg.InitialInventory,
		}

		t.Run(test.name, func(t *testing.T) {

			produce, err := GS.GetAllProduce()

			if (err != nil) != test.errExpected {
				t.Errorf("errExpected: %v, errReceived: %v", test.errExpected, (err != nil))
			}

			if (produce != nil) != test.returnExpected {
				t.Errorf("returnExpected: %v, returnReceived: %v", test.returnExpected, (produce != nil))
			}

			if produce != nil {
				if test.lenExpected != len(GS.inventory) {
					t.Errorf("lenExpected: %d, lenFound: %d", test.lenExpected, len(GS.inventory))
				}
			}
		})
	}
}
