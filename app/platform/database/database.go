package database

import (
	"fiber-api-example/app/models/books"
	"fiber-api-example/app/utils/logger"
	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

// Queries struct for collect all app queries.
type Queries struct {
	*books.BookQueries // load queries from Book model
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Username string
	Password string
	Port     int
	Database string
}

func (cfg *DatabaseConfig) ConnectString() string {
	s := []string{cfg.Username, ":", cfg.Password, "@", cfg.Host, ":", strconv.Itoa(cfg.Port), "/", cfg.Database}
	return strings.Join(s, "")
}

func Connection() (*Queries, error) {
	config := DatabaseConfig{
		Driver:   "postgres",
		Host:     viper.GetString("DB"),
		Username: viper.GetString("DB_USERNAME"),
		Password: viper.GetString("DB_PASSWORD"),
		Port:     viper.GetInt("DB_PORT"),
		Database: viper.GetString("DB_DATABASE"),
	}
	db, err := sqlx.Open("pgx", config.ConnectString())
	if err != nil {
		logger.Error(err, "Can't connect to "+config.Database)
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 5)

	return &Queries{
		// Set queries from models:
		BookQueries: &books.BookQueries{DB: db}, // from Book model
	}, nil
}
