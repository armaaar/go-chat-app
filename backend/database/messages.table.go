package database

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Name    string
	Message string
}
