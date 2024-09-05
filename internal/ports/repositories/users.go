package ports_repositories

import "github.com/Arthur-Conti/keepass-project/internal/models"

type UsersRepositoryInterface interface {
	ListAllUsers() ([]models.User, error)
	GetUserByID(string) (models.User, error)
	CreateUser(models.UserModel) error
	UpdateUser(string)
}
