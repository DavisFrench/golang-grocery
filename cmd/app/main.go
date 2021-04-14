package main

import (
	"fmt"

	gg "DavisFrench/golang-grocery"
	"DavisFrench/golang-grocery/db"
)

type App struct {
	groceryService gg.GroceryService
}

func main() {
	groceryService := db.NewGroceryService()

	app := App{
		groceryService: groceryService,
	}

	app.run()
}

func (a *App) run() error {
	fmt.Println("main")

	return nil
}
