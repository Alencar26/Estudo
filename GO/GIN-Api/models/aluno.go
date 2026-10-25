package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11,regexp=^[0-9]{11}$"`
	RG   string `json:"rg" validate:"len=9,regexp=^[0-9]{8}[0-9X]$"`
}

func (a *Aluno) Validate() error {
	if err := validator.Validate(a); err != nil {
		return err
	}
	return nil
}
