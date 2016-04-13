package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Session struct {
	gorm.Model
	UserID            uint
	Cookie            string
	OriginallyCreated time.Time
	LastUsed          time.Time
}

func FromCookie(cookie string) (session Session) {
	return
}

func ActiveSession() {

}
