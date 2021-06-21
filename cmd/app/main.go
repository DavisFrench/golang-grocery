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

func envGet(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("env not found")
	}
	return value
}

var (
	host = envGet("PGHOST")
	port = envGet("PGPORT")
	user = envGet("PGUSER")
	password = envGet("PGPASSWORD")
	dbname = envGet("PGDATABASE")
)

func main() {

	// groceryService := db.NewGroceryService()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	groceryService, err := db.NewPgService(psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

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
