package config

import (
	"log"
	"os"
	"github.com/go-pg/pg/v9"

	controllers "createrestful/controllers"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User: "postgres",
		Password: "1150",
		Addr: "localhost:5432",
		Database: "golangpostgres",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)

	return db
}