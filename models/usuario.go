package models

import (
	"github.com/jinzhu/gorm"
)

// Usuario representa uma conta
type Usuario struct {
	gorm.Model

	Nome string
	Email string
	Senha []byte
}
