package validator

import (
	"explorer-api/models/validator"
	"github.com/jinzhu/gorm"
)

func GetByPubKey(db *gorm.DB, pubKey string) validator.Validator {
	var v validator.Validator
	db.Where("pub_key = ?", pubKey).First(&v)
	return v
}
