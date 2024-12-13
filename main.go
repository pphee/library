package main

import (
	"github.com/joho/godotenv"
	"github.com/pphee/library/handler"
	"github.com/pphee/library/repository"
	"github.com/pphee/library/store"
	"github.com/pphee/library/use_case"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Please provide a DATABASE_URL in the environment or .env file")
	}
	dbStore := store.NewPostgresStore(dsn)

	r := gin.Default()

	Repo := repository.New(dbStore.DB)
	useCase := use_case.New(Repo)
	bookHandler := handler.New(useCase)

	bookHandler.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
