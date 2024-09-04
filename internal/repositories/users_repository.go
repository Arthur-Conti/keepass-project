package repositories

import (
	"github.com/Arthur-Conti/keepass-project/internal/config"
	"github.com/Arthur-Conti/keepass-project/internal/models"
)

type UsersRepository struct {}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (repository *UsersRepository) ListAllUsers() {

}

func (repository *UsersRepository) GetUserByID(id string) {

}

func (repository *UsersRepository) CreateUser(user models.UserModel) error {
	postgres := config.NewPostgresConnection()
	databaseConnection, context, err := postgres.Connect()
	if err != nil {
		return err
	}
	defer postgres.CloseConnection(databaseConnection, context)
	createUserQuery := `
	INSERT INTO users (id, first_name, last_name, email, password)
	VALUES ($1, $2, $3, $4, $5)
	`
	postgres.RunQuery(databaseConnection, context, createUserQuery, user.ID, user.FirstName, user.LastName, user.Email, user.Password)
	return nil
}

func (repository *UsersRepository) UpdateUser(id string) {

}