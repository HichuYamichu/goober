package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Service performs operations specyfic to user domain
type Service struct {
	usrRepo *Repository
}

// NewService creates new user service
func NewService(usrRepo *Repository) *Service {
	s := &Service{usrRepo: usrRepo}
	return s
}

// CreateUser creates a user
func (s Service) CreateUser(user *User) error {
	hash, err := s.hashPassword(user.Pass)
	if err != nil {
		return err
	}
	user.Pass = hash
	return s.usrRepo.Create(user)
}

// DeleteUser deletes a user
func (s Service) DeleteUser(id int) error {
	return s.usrRepo.Delete(id)
}

// ChangePassword changes user's password
func (s Service) ChangePassword(userID int, pass string) error {
	where := &User{ID: userID}
	hash, err := s.hashPassword(pass)
	if err != nil {
		return err
	}
	fields := &User{Pass: hash}
	return s.usrRepo.Update(where, fields)
}

// VerifyCredentials verifies user credentials
func (s Service) VerifyCredentials(username, password string) (*User, error) {
	user, err := s.findOneByUsername(username)
	if err != nil {
		return nil, err
	}

	match := s.checkPasswordHash(password, user.Pass)
	if !match {
		return nil, errors.New("wrong password")
	}

	return user, nil
}

func (s Service) findOneByUsername(username string) (*User, error) {
	user := &User{Username: username}
	user, err := s.usrRepo.FindOne(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s Service) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s Service) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
