package repository

import (
	"BookStore/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByPhone(phone string) (*domain.User, error)
	FindByID(id uint) (*domain.User, error)
	Update(user *domain.User) error
	Delete(user *domain.User) error
	ListAll() ([]domain.User, error)
	SearchByName(name string) ([]domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByPhone(phone string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(user *domain.User) error {
	return r.db.Delete(user).Error
}

func (r *userRepository) ListAll() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) SearchByName(name string) ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Where("name LIKE ?", name).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
