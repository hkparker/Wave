package models

import (
	"github.com/jinzhu/gorm"
)

var db gorm.DB

func SetDB(wave_db gorm.DB) {
	db = wave_db
}
