package models

import (
	"github.com/jinzhu/gorm"
)

type Alert struct {
	gorm.Model
	Title    string
	Rule     string
	Severity string
}

func (alert *Alert) Save() error {
	return Orm.Save(&alert).Error
}

func (alert *Alert) Delete() error {
	return Orm.Delete(&alert).Error
}
