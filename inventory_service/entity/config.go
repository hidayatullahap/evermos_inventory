package entity

import (
	"os"

	"github.com/hidayatullahap/evermos/pkg/db/mysql"
)

type Config struct {
	mysql.MysqlConnection
}

func NewConfig() Config {
	var c Config

	sql := mysql.MysqlConnection{
		Username: os.Getenv("DB_USERNAME"),
		Pass:     os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
	}

	c.MysqlConnection = sql
	return c
}
