package main

import "gorm.io/gorm"

type KeyValuePairModel struct {
	gorm.Model
	Key   string
	Value string
	Ttl   int
}
type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Ttl   int    `json:"ttl"`
}
