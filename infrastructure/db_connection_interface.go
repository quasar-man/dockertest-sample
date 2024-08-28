package infrastructure

import (
	"gorm.io/gorm"
)

type DBConnectInterface interface {
	DsnString() string
	DBConnectionOpen() (*gorm.DB, error)
	GormAutoMigrate(*gorm.DB)
}
