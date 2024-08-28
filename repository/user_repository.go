package repository

import (
	"gorm.io/gorm"
	"github.com/quasar-man/dockertest-sample/entity"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (u *UserRepository) FindAll() (*[]entity.User, error) {
	users := []entity.User{}
	u.db.Find(&users)

	return &users, nil
}

func (u *UserRepository) FindByID(id uint) (*entity.User, error) {
	user := entity.User{}
	u.db.First(&user, id)

	return &user, nil
}

func (u *UserRepository) FindByEmail(email string) (*entity.User, error) {
	user := entity.User{}
	u.db.Where("email = ?", email).First(&user)

	return &user, nil
}
