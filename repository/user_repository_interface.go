package repository

import (
	"github.com/quasar-man/dockertest-sample/entity"
)

type UserRepositoryInterface interface {
	FindAll() (*[]entity.User, error)
	FindByID(id uint) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}
