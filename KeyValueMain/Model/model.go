package model

import "gorm.io/gorm"

type KeyValuePairModel struct {
	gorm.Model
	Key   string
	Value string
	Ttl   int
}

// gorm.Model equals
// type User struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// 	Name string
//   }
