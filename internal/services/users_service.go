package services

import (
	"github.com/Arthur-Conti/keepass-project/internal/config"
	"github.com/Arthur-Conti/keepass-project/internal/models"
	ports_repositories "github.com/Arthur-Conti/keepass-project/internal/ports/repositories"
)

var logger = config.GetLogger()

type usersService struct {
	repository ports_repositories.UsersRepositoryInterface
}

func NewUsersService(repository ports_repositories.UsersRepositoryInterface) *usersService {
	return &usersService{
		repository: repository,
	}
}

func (service *usersService) ListAllUsers() ([]models.User, error) {
	return service.repository.ListAllUsers()
}

func (service *usersService) GetUserByID(id string) (models.User, error) {
	return service.repository.GetUserByID(id)
}

func (service *usersService) CreateUser(user models.UserModel) error {
	userID, err := config.CreateID()
	if err != nil {
		logger.Error("error creating user id: " + err.Error())
		return err
	}
	user.ID = userID
	return service.repository.CreateUser(user)
}

func (service *usersService) UpdateUser(id string) {

}
