package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ConnectionInfo struct {
	Host     string
	User     string
	Password string
	Port     int
	DBName   string
}

type ConnectorFunc func() (*sqlx.DB, error)

func NewMySQLConnection(info ConnectionInfo) (*sqlx.DB, error) {
	connectionStr := fmt.Sprintf("%s:%s@(%s:%d)/%s", info.User, info.Password, info.Host, info.Port, info.DBName)
	db, err := sqlx.Open("mysql", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("unable to create mysql connection: %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection flew away: %s", err.Error())
	}

	return db, nil
}
