package services

import (
	"github.com/Arthur-Conti/keepass-project/internal/config"
	"github.com/Arthur-Conti/keepass-project/internal/models"
	"github.com/Arthur-Conti/keepass-project/internal/repositories"
)

type usersService struct {
	repository *repositories.UsersRepository
}

func NewUsersService() *usersService {
	return &usersService{
		repository: repositories.NewUsersRepository(),
	}
}

func (service *usersService) ListAllUsers() {

}

func (service *usersService) GetUserByID(id string) {

}

func (service *usersService) CreateUser(user models.UserModel) (error) {
	userID, err := config.CreateID()
	if err != nil {
		return err
	}
	user.ID = userID
	return service.repository.CreateUser(user) 
}

func (service *usersService) UpdateUser(id string) {

}