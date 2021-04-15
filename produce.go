package grocery

type Produce struct {
	Name        string  `json:"name"`
	ProduceCode string  `json:"produce_code"`
	UnitPrice   float32 `json:"unit_price"`
}

// defines all the methods expected for anything implementing the GroceryService
type GroceryService interface {
	AddProduce(produce Produce) error
	DeleteProduce(produceCode string) error
	GetProduceByCode(produceCode string) (*Produce, error)
	GetAllProduce() ([]Produce, error)
}

var InitialInventory = []Produce{
	Produce{
		ProduceCode: "A12T-4GH7-QPL9-3N4M",
		Name:        "Lettuce",
		UnitPrice:   3.46,
	},
	Produce{
		ProduceCode: "E5T6-9U13-TH15-QR88",
		Name:        "Peach",
		UnitPrice:   2.99,
	},
	Produce{
		ProduceCode: "YRT6-72AS-K736-L4AR",
		Name:        "Green Pepper",
		UnitPrice:   0.79,
	},
	Produce{
		ProduceCode: "TQ4C-VV6T-75ZX-1RMR",
		Name:        "Gala Apple", //honey crisp is where it is at
		UnitPrice:   3.59,
	},
}
