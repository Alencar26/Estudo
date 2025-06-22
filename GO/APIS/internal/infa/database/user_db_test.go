package database

import (
	"testing"

	"github.com/Alencar26/Estudo/GO/APIS/internal/entity"
	"github.com/Alencar26/Estudo/GO/APIS/pkg/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createUserForTest() (*entity.User, error) {
	letters := test.RandomLetter(4)
	userTest, err := entity.NewUser("user"+letters+"@mail.com", "UserName"+letters, "senha123")
	if err != nil {
		return nil, err
	}
	return userTest, nil
}

func TestCreateUser(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	userTest, err := createUserForTest()
	assert.Nil(t, err)
	assert.NotNil(t, userTest)

	userDB := NewUser(db)

	err = userDB.Create(*userTest)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", userTest.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, userTest.Name, userFound.Name, "Nome do usuário é diferente do esperado")
	assert.Equal(t, userTest.Email, userFound.Email, "E-mail do usuário é diferente do esperado")
	assert.Equal(t, userTest.ID, userFound.ID, "ID do usuário é diferente do esperado")
	assert.NotNil(t, userFound.Password)
}

func TestFindUserByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	userTest, err := createUserForTest()
	assert.Nil(t, err)
	assert.NotNil(t, userTest)

	userDB := NewUser(db)

	err = userDB.Create(*userTest)

	userFound, err := userDB.FindByEmail(userTest.Email)

	assert.Nil(t, err)
	assert.Equal(t, userTest.Name, userFound.Name, "Nome do usuário é diferente do esperado")
	assert.Equal(t, userTest.Email, userFound.Email, "E-mail do usuário é diferente do esperado")
	assert.Equal(t, userTest.ID, userFound.ID, "ID do usuário é diferente do esperado")
	assert.NotNil(t, userFound.Password)
}
