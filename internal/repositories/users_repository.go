package repositories

import (
	"github.com/Arthur-Conti/keepass-project/internal/config"
	"github.com/Arthur-Conti/keepass-project/internal/models"
)

var logger = config.GetLogger()

type UsersRepository struct{}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (repository *UsersRepository) ListAllUsers() ([]models.User, error) {
	postgres := config.NewPostgresConnection()
	err := postgres.Connect()
	if err != nil {
		logger.Error("error connecting to postgres: " + err.Error())
		return nil, err
	}
	defer postgres.CloseConnection()
	listAllUsersQuery := `SELECT id, first_name, last_name, email, created_time FROM users`
	var users []models.User
	rows, err := postgres.RunQuery(listAllUsersQuery)
	if err != nil {
		logger.Error("error listing all users: " + err.Error())
		return nil, err
	}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedTime)
		if err != nil {
			logger.Error("error scannig rows: " + err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository *UsersRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	postgres := config.NewPostgresConnection()
	err := postgres.Connect()
	if err != nil {
		logger.Error("error connecting to postgres: " + err.Error())
		return user, err
	}
	defer postgres.CloseConnection()
	logger.Info(id)
	getUserByIDQuery := `SELECT id, first_name, last_name, email, created_time FROM users WHERE id = $1`
	err = postgres.RunQueryRow(getUserByIDQuery, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedTime)
	if err != nil {
		logger.Error("error scannig rows: " + err.Error())
		return user, err
	}
	return user, nil
}

func (repository *UsersRepository) CreateUser(user models.UserModel) error {
	postgres := config.NewPostgresConnection()
	err := postgres.Connect()
	if err != nil {
		logger.Error("error connecting to postgres: " + err.Error())
		return err
	}
	defer postgres.CloseConnection()
	createUserQuery := `
	INSERT INTO users (id, first_name, last_name, email, password)
	VALUES ($1, $2, $3, $4, $5)
	`
	postgres.RunQueryRow(createUserQuery, user.ID, user.FirstName, user.LastName, user.Email, user.Password)
	return nil
}

func (repository *UsersRepository) UpdateUser(id string) {

}
