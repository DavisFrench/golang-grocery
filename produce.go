package grocery

type Produce struct {
	Name        string  `json:"name"`
	ProduceCode string  `json:"produce_code"`
	UnitPrice   float32 `json:"unit_price"`
}

type GroceryService interface {
	AddProduce(produce Produce) error
	DeleteProduce(produceCode string) error
	GetProduceByCode(produceCode string) (*Produce, error)
	GetAllProduce() ([]Produce, error)
}
