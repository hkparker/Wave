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
	var config TLS
	err = database.Orm.First(&config).Error
	if err != nil {

	}
	config.CaCert = []byte(request["ca_cert"])
	config.CaCert = []byte(request["ca_cert"])
	err = database.Orm.Save(&config).Error
	if err != nil {

	}
	return
}

func createTLSIfMissing() {}
