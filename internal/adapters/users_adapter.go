package adapters

import (
	"github.com/Arthur-Conti/keepass-project/internal/ports/services"
	"github.com/Arthur-Conti/keepass-project/internal/repositories"
	"github.com/Arthur-Conti/keepass-project/internal/services"
)

func NewUserServiceAdapter() ports_services.UsersServiceInterface {
	userRepository := repositories.NewUsersRepository()
	userService := services.NewUsersService(userRepository)
	return userService
}