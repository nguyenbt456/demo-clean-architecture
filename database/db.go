package database

import (
	"fmt"

	"github.com/nguyenbt456/demo-clean-architecture/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	POSTGRES_HOST     = "POSTGRES_HOST"
	POSTGRES_PORT     = "POSTGRES_PORT"
	POSTGRES_USER     = "POSTGRES_USER"
	POSTGRES_PASSWORD = "POSTGRES_PASSWORD"
	POSTGRES_DBNAME   = "POSTGRES_DBNAME"
	POSTGRES_SSLMODE  = "POSTGRES_SSLMODE"

	pgHost     = "127.0.0.1"
	pgPort     = "5432"
	pgUser     = "postgres"
	pgPassword = "123456"
	pgDBName   = "demo"
	pgSSLMode  = "disable"
)

type postgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func ConnectDB() {
	cfg := postgresConfig{
		Host:     util.EVString(POSTGRES_HOST, pgHost),
		Port:     util.EVString(POSTGRES_PORT, pgPort),
		User:     util.EVString(POSTGRES_USER, pgUser),
		Password: util.EVString(POSTGRES_PASSWORD, pgPassword),
		DBName:   util.EVString(POSTGRES_DBNAME, pgDBName),
		SSLMode:  util.EVString(POSTGRES_SSLMODE, pgSSLMode),
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	var err error
	if db, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}), nil); err != nil {
		panic(err)
	}
}
