package ports_services

import "github.com/Arthur-Conti/keepass-project/internal/models"

type UsersServiceInterface interface {
	ListAllUsers() ([]models.User, error)
	GetUserByID(string) (models.User, error)
	CreateUser(models.UserModel) error
	UpdateUser(string)
}