package main

import (
	"fmt"
	"github.com/mm-go-rest-api/internal/comment"
	"github.com/mm-go-rest-api/internal/db"
	transportHttp "github.com/mm-go-rest-api/internal/transport/http"
)

func Run() error {
	fmt.Println("Starting up the application...")
	database,err := db.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := database.MigrateBD(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}
	fmt.Println("Successfully connected to the database")

	cmtService := comment.NewService(database)

	httpHandler := transportHttp.NewHandler(cmtService)

	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main()  {
	fmt.Println("EntryPoint: Rest API")

	if err:= Run(); err != nil {
		fmt.Println(err)
	}
}
