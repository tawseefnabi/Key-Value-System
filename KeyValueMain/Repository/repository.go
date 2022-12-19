package repository

import (
	model "KeyValueMain/Model"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(dbName string) *Repository {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	// Migrate the schema
	fmt.Println("Connected to db")
	db.AutoMigrate(&model.KeyValuePairModel{})
	return &Repository{
		Db: db,
	}
}
