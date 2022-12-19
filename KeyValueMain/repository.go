package main

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(dbName string) *Repository {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&KeyValuePairModel{})
	return &Repository{
		db: db,
	}
}

func (r *Repository) put(keyValuePair KeyValuePair) error {
	var keyValuePairModel KeyValuePairModel
	r.db.Where("key = ?", keyValuePair.Key).First(&keyValuePairModel)
	if keyValuePairModel.Key != keyValuePair.Key {
		r.db.Create(&KeyValuePairModel{
			Key:   keyValuePair.Key,
			Value: keyValuePair.Value,
			Ttl:   keyValuePair.Ttl,
		})
	} else {
		keyValuePairModel.Value = keyValuePair.Value
		keyValuePairModel.Ttl = keyValuePair.Ttl
		r.db.Save((&keyValuePairModel))
	}
	return nil

}
func (r *Repository) get(key string) (KeyValuePairModel, error) {
	var keyValuePairModel KeyValuePairModel
	r.db.Where("key = ?", key).First(&keyValuePairModel)
	if keyValuePairModel.Key == "" {
		return keyValuePairModel, errors.New("Key not present in the system")
	}
	return keyValuePairModel, nil
}
