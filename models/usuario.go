package models

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/validator.v2"
	"golang.org/x/crypto/bcrypt"
)

// Usuario representa uma conta
type Usuario struct {
	gorm.Model

	Nome string `gorm:"not null;type:varchar(40)" validate:"min=3,max=40,regexp=^[a-zA-Z0-9_-]*$"`
	Email string `gorm:"not null;type:varchar(100);unique" validate:"regexp=^[A-Z0-9a-z._%+-]{3,}+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,4}$"`
	Senha string `gorm:"not null;type:varchar(120)"`
}

func (u *Usuario)SetHashedSenha(senha string) {

}

func NewUsuario(nome, email, senha string) (*Usuario, error) {

	hashedSenha, errs := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if errs != nil {
		return nil, errs
	}

	u := &Usuario{Nome:nome, Email:email, Senha:string(hashedSenha)}

	if errs = validator.Validate(u); errs != nil {
		return nil, errs
	}

	return u, nil
}