package repository

import (
	"github.com/loak155/techbranch-backend/internal/user/domain"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *domain.User) error
	GetUser(id int) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	ListUsers() (*[]domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user *domain.User) error {
	err := ur.db.Create(user).Error
	return err
}

func (ur *userRepository) GetUser(id int) (*domain.User, error) {
	user := &domain.User{}
	err := ur.db.First(user, id).Error
	return user, err
}

func (ur *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	err := ur.db.Where("email=?", email).First(user).Error
	return user, err
}

func (ur *userRepository) ListUsers() (*[]domain.User, error) {
	var users *[]domain.User
	err := ur.db.Find(&users).Error
	return users, err
}

func (ur *userRepository) UpdateUser(user *domain.User) error {
	err := ur.db.Save(user).Error
	return err
}

func (ur *userRepository) DeleteUser(id int) error {
	err := ur.db.Delete(&domain.User{}, id).Error
	return err
}
