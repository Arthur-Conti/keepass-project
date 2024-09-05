package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type postgresConnection struct {
	Username     string
	Password     string
	IP           string
	DatabaseName string
	Context      context.Context
	Connection   *pgx.Conn
}

func NewPostgresConnection() *postgresConnection {
	return &postgresConnection{
		Username:     GetEnv("POSTGRES_USER"),
		Password:     GetEnv("POSTGRES_PASSWORD"),
		IP:           GetEnv("POSTGRES_IP"),
		DatabaseName: GetEnv("POSTGRES_DB"),
	}
}

func (postgres *postgresConnection) Connect() error {
	err := postgres.CreateConnection()
	if err != nil {
		return err
	}
	return nil
}

func (postgres *postgresConnection) CreateConnection() error {
	connectionString := fmt.Sprintf("postgres://%v:%v@%v/%v", postgres.Username, postgres.Password, postgres.IP, postgres.DatabaseName)
	context := context.Background()
	connection, err := pgx.Connect(context, connectionString)
	if err != nil {
		return err
	}
	postgres.Context = context
	postgres.Connection = connection
	return nil
}

func (postgres *postgresConnection) CloseConnection() {
	postgres.Connection.Close(postgres.Context)
}

func (postgres *postgresConnection) RunQuery(query string, queryParams ...any) (pgx.Rows, error) {
	return postgres.Connection.Query(postgres.Context, query, queryParams...)
}

func (postgres *postgresConnection) RunQueryRow(query string, queryParams ...any) pgx.Row {
	return postgres.Connection.QueryRow(postgres.Context, query, queryParams...)
}
