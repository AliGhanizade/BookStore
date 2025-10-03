package service

import (
	"BookStore/internal/domain"
	"BookStore/internal/repository"
	"errors"
)

type UserService interface {
	Register(name, phone, password string) (*domain.User, error)
	Login(phone, password string) (*domain.User, error)
	GetByID(id uint) (*domain.User, error)
	Update(user *domain.User) error
	Delete(user *domain.User) error
	ListAll() ([]domain.User, error)
	SearchByName(name string) ([]domain.User, error)

}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) Register(name, phone, password string) (*domain.User, error) {
	user := &domain.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	return user, s.repo.Create(user)
}

func (s *userService) Login(phone, password string) (*domain.User, error) {
	user, err := s.repo.FindByPhone(phone)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (s *userService) GetByID(id uint) (*domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) Update(user *domain.User) error {
	return s.repo.Update(user)
}

func (s *userService) Delete(user *domain.User) error {
	return s.repo.Delete(user)
}

func (s *userService) ListAll() ([]domain.User, error) {
	return s.repo.ListAll()
}

func (s *userService) SearchByName(name string) ([]domain.User, error){
	return s.repo.SearchByName(name)
}