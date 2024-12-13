package store

import (
	modelgorm "github.com/pphee/library/store/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PostgresStore struct {
	DB *gorm.DB
}

func NewPostgresStore(dsn string) *PostgresStore {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error getting database connection: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	if err := db.AutoMigrate(&modelgorm.Book{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	return &PostgresStore{
		DB: db,
	}
}
