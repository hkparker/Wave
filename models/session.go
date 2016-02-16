package models

import (
	"github.com/jinzhu/gorm"
)

type Session struct {
	gorm.Model
	Cookie string
}

func CreateNewSession() {}
