package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbName = "../keyValue.db"
)

type KeyValuePairModel struct {
	gorm.Model
	Key   string
	Value string
	Ttl   int
}

func main() {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	fmt.Println("db", db)
	if err != nil {
		panic("failed to connect database")
	}
	s := gocron.NewScheduler(time.UTC)
	cnt := 1
	s.Every(5).Seconds().Do(func() {
		fmt.Println("cron performing operation: ", cnt)
		var keyValuePairModel []KeyValuePairModel
		db.Where("ttl> ? or ttl=-1", time.Now().Unix()).Find(&keyValuePairModel)
		keyList := []int{}
		for _, keyValue := range keyValuePairModel {
			id := keyValue.ID
			keyList = append(keyList, int(id))
		}
		if len(keyList) > 0 {
			fmt.Println("Below id are deleted:")
			fmt.Println(keyList)
			db.Delete(&KeyValuePairModel{}, keyList)
		}
		cnt++
	})
	s.StartBlocking()
	// s.StartAsync()
}
