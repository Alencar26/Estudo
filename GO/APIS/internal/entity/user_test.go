package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	user, err := NewUser("email@email.com", "Nome Genérico", "Senha123")
	assert.Nil(t, err)
	assert.NotNil(t, user, "User não pode ser nulo")

	assert.NotEmpty(t, user.ID, "ID não pode ser vazio")
	assert.NotEmpty(t, user.Email, "Email não pode ser vazio")
	assert.NotEmpty(t, user.Name, "Nome não pode ser vazio")
	assert.NotEmpty(t, user.Password, "Senha não pode ser vazia")

	assert.Equal(t, "Nome Genérico", user.Name, "Nome deve ser igual a 'Nome Genérico'")
	assert.Equal(t, "email@email.com", user.Email, "Email deve ser igual a 'email@email.com'")
}

func TestValidadePassword(t *testing.T) {
	user, err := NewUser("email@email.com", "Nome Genérico", "SenhaDiferente12345")
	assert.Nil(t, err)
	assert.NotNil(t, user, "User não pode ser nulo")

	assert.NotEqual(t, "SenhaDiferente12345", user.Password, "Senha deve ser diferente de 'SenhaDiferente12345'")
	assert.True(t, user.ValidatePassword("SenhaDiferente12345"), "Deveria ser verdadeiro e recebeu falso")
	assert.False(t, user.ValidatePassword("senhaDiferente123456"), "Deveria ser falso e recebeu verdadeiro")
}
