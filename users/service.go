package users

import (
	"github.com/hichuyamichu-me/uploader/errors"
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
	const op errors.Op = "users/service.CreateUser"

	hash, err := s.hashPassword(user.Pass)
	if err != nil {
		return errors.E(err, errors.Internal, op)
	}
	user.Pass = hash
	return s.usrRepo.Create(user)
}

// DeleteUser deletes a user
func (s Service) DeleteUser(id int) error {
	const op errors.Op = "users/service.DeleteUser"

	return s.usrRepo.Delete(id)
}

// ChangePassword changes user's password
func (s Service) ChangePassword(userID int, pass string) error {
	const op errors.Op = "users/service.ChangePassword"

	where := &User{ID: userID}
	hash, err := s.hashPassword(pass)
	if err != nil {
		return errors.E(err, errors.Internal, op)
	}
	fields := &User{Pass: hash}
	return s.usrRepo.Update(where, fields)
}

// VerifyCredentials verifies user credentials
func (s Service) VerifyCredentials(username, password string) (*User, error) {
	const op errors.Op = "users/service.VerifyCredentials"

	user, err := s.findOneByUsername(username)
	if err != nil {
		return nil, errors.E(err, errors.Authentication, op)
	}

	match := s.checkPasswordHash(password, user.Pass)
	if !match {
		return nil, errors.E(err, errors.Authentication, op)
	}

	return user, nil
}

func (s Service) findOneByUsername(username string) (*User, error) {
	const op errors.Op = "users/service.findOneByUsername"

	user := &User{Username: username}
	user, err := s.usrRepo.FindOne(user)
	if err != nil {
		return nil, errors.E(err, op)
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
