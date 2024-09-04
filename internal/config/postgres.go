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
}

func NewPostgresConnection() *postgresConnection {
	return &postgresConnection{
		Username:     GetEnv("POSTGRES_USER"),
		Password:     GetEnv("POSTGRES_PASSWORD"),
		IP:           GetEnv("POSTGRES_IP"),
		DatabaseName: GetEnv("POSTGRES_DB"),
	}
}

func (postgres *postgresConnection) Connect() (*pgx.Conn, context.Context, error) {
	connectionString := fmt.Sprintf("postgres://%v:%v@%v/%v", postgres.Username, postgres.Password, postgres.IP, postgres.DatabaseName)
	context := context.Background()
	connection, err := pgx.Connect(context, connectionString)
	if err != nil {
		return nil, nil, err
	}
	return connection, context, nil
}

func (postgres *postgresConnection) CloseConnection(connection *pgx.Conn, context context.Context) {
	connection.Close(context)
}

func (postgres *postgresConnection) RunQuery(connection *pgx.Conn, context context.Context, query string, queryParams ...any) pgx.Row {
	return connection.QueryRow(context, query, queryParams...)
}
