package tests

import (
	"testing"
	"github.com/Arthur-Conti/keepass-project/internal/models"
	"github.com/Arthur-Conti/keepass-project/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	settingUpEnvs()
	repository := repositories.NewUsersRepository()
	err := repository.CreateUser(models.UserModel{
		ID: "1234",
		FirstName: "test",
		LastName: "user",
		Email: "user_test@test.com",
		Password: "testpass123",
	})
	assert.NoError(t, err)
}

func TestListAllUsers(t *testing.T) {
	settingUpEnvs()
	repository := repositories.NewUsersRepository()
	users, err := repository.ListAllUsers()
	if assert.NoError(t, err) {
		_, ok := interface{}(users[0]).(models.User)
		assert.False(t, ok, "users must be of type models.User")
	}
}
