package main

import (
	"context"
	"log"
	"os"

	"codeid.northwind/config"
	"codeid.northwind/repositories"
	"codeid.northwind/server"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("starting northwind restapi")

	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)
	log.Println(dbHandler)

	// test insert to category, using goroutine
	ctx := context.Background()
	queries := repositories.New(dbHandler)

	newCategory, err := queries.CreateCategory(ctx,
		repositories.CreateCategoryParams{
			CategoryID:   101,
			CategoryName: "Mainan",
			Description:  "Mainan anak",
			Picture:      nil,
		},
	)

	if err != nil {
		log.Fatal("Error : ", err)
	}
	log.Println(newCategory)

}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "northwind" + env
	}

	return "northwind"
}
