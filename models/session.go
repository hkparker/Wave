package models

import (
	"github.com/jinzhu/gorm"
)

type Session struct {
	gorm.Model
	Cookie   string
	LastUsed string
}

func CreateNewSession() {}
