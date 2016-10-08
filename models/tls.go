package models

import (
	"github.com/hkparker/Wave/database"
	"github.com/jinzhu/gorm"
)

type TLS struct {
	gorm.Model
	CaCert     []byte
	PrivateKey []byte
}

func TLSConfig() (config TLS, err error) {
	createTLSIfMissing()
	err = database.Orm.First(&config).Error
	if err != nil {

	}
	return
}

func SetTLS(request map[string]string) (err error) {
	createTLSIfMissing()
	// ensure no collectors
	var config TLS
	err = database.Orm.First(&config).Error
	if err != nil {

	}
	config.CaCert = []byte(request["ca_cert"])
	config.PrivateKey = []byte(request["private_key"])
	err = database.Orm.Save(&config).Error
	if err != nil {

	}
	return
}

func createTLSIfMissing() (err error) {
	var count int
	var tls []TLS
	err = database.Orm.Find(&tls).Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {

	}
	return
}
