package infrastructure

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/quasar-man/dockertest-sample/entity"
)

type MySQLConnection struct {}

func NewMySQLConnection() DBConnectInterface {
	return &MySQLConnection{}
}

func (m *MySQLConnection) DsnString() string {
	return "user:password@tcp(db:3306)/dockertest_db?charset=utf8mb4&parseTime=True&loc=Local"
}

func (m *MySQLConnection) DBConnectionOpen() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(m.DsnString()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (m *MySQLConnection) GormAutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}
