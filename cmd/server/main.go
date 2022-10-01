package main

import (
	"fmt"
	"log"

	"github.com/by-sabbir/company-microservice-rest/pkg/db"
)

func Run() error {
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}
	log.Println("Successfully Connected to the DB!")

	if err := db.MigrateDB(); err != nil {
		return fmt.Errorf("migrations failed because of: %w", err)
	}

	// cmpService := company.NewService(db)
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatalf("could not run the service: %+v\n", err)
	}
}
