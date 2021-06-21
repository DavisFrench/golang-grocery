package db

import (
	gg "DavisFrench/golang-grocery"
	"database/sql"

	_ "github.com/lib/pq"
)

type pgService struct {
	db *sql.DB
}

func NewPgService(psqlInfo string) (*pgService, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return &pgService{
		db: db,
	}, nil
}

func (pg *pgService) close() error {
	return pg.db.Close()
}

func (gs *pgService) AddProduce(produce gg.Produce) error {
	return nil
}

func (gs *pgService) DeleteProduce(produceCode string) error {
	return nil
}

func (gs *pgService) GetProduceByCode(produceCode string) (*gg.Produce, error) {
	return nil, nil
}

func (pg *pgService) GetAllProduce() ([]gg.Produce, error) {
	rows, err := pg.db.Query("SELECT * FROM inventory")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var inventory []gg.Produce
	for rows.Next() {
		var produce gg.Produce
		err = rows.Scan(&produce.ProduceCode, &produce.Name, &produce.UnitPrice)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return inventory, nil
}