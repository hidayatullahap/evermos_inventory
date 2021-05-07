package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConnection struct {
	Username string
	Pass     string
	Host     string
	Port     string
	DbName   string
}

// Connect create connection for gorm
func Connect(conf MysqlConnection) (*gorm.DB, error) {
	con := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", conf.Username, conf.Pass, conf.Host, conf.Port, conf.DbName)
	return gorm.Open(mysql.Open(con), &gorm.Config{})
}
