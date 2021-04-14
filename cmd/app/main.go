package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	gg "DavisFrench/golang-grocery"
	"DavisFrench/golang-grocery/db"
	"DavisFrench/golang-grocery/http"
)

type App struct {
	groceryService gg.GroceryService
}

func main() {
	groceryService := db.NewGroceryService()

	app := App{
		groceryService: groceryService,
	}

	if err := app.run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Shutting down from interrupt")
}

func (a *App) run() error {

	httpServer := http.NewServer(a.groceryService)
	httpServer.Addr = ":8888"

	if err := httpServer.Open(); err != nil {
		return err
	}

	fmt.Println("Listening on " + httpServer.Addr)

	return nil
}
