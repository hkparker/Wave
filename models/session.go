package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Session struct {
	gorm.Model
	UserID   uint
	Cookie   string
	LastUsed time.Time
}
