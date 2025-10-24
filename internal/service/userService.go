package service

import (
	"errors"
	"fmt"
	"go-ecommerce/internal/domain"
	"go-ecommerce/internal/dto"
	"go-ecommerce/internal/repository"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s UserService) SignUp(input dto.UserSignup) (string, error) {
	log.Println(input)

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})

	// generate token
	log.Println(user)

	userInfo := fmt.Sprintf("user id: %d, email: %s", user.ID, user.Email)

	return userInfo, err
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	// perform some db operation
	// business logic

	user, err := s.Repo.FindUser(email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s UserService) Login(email, password string) (string, error) {
	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user doest not exist with the provided email id")
	}

	// compare password and generate token

	return user.Email, nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId uint) ([]interface{}, error) {
	return nil, nil
}
