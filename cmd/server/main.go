package main

import (
	"fmt"
	"github.com/mm-go-rest-api/internal/db"
)

func Run() error {
	fmt.Println("Starting up the application...")
	db,err := db.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateBD(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}
	fmt.Println("Successfully connected to the database")
	fmt.Println(db)
	return nil
}

func main()  {
	fmt.Println("EntryPoint: Rest API")

	if err:= Run(); err != nil {
		fmt.Println(err)
	}
}
