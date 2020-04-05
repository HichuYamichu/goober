package users

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/hichuyamichu-me/goober/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
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

// ChangePassword changes user's password
func (s *Service) ChangePassword(userID uuid.UUID, pass string) error {
	const op errors.Op = "users/service.ChangePassword"

	hash, err := s.HashPassword(pass)
	if err != nil {
		return errors.E(err, errors.Internal, op)
	}

	user := &User{ID: userID, Pass: hash}
	return s.usrRepo.Update(user)
}

// CreateUser creates user
func (s *Service) CreateUser(username string, password string) error {
	const op errors.Op = "users/service.CreateUser"

	quota := viper.GetInt64("quota")
	password, err := s.HashPassword(password)
	if err != nil {
		return errors.E(err, errors.Internal, op)
	}

	token, err := s.GenerateToken(username)
	if err != nil {
		return errors.E(err, op)
	}

	user := &User{Username: username, Pass: password, Admin: false, Active: false, Quota: quota, Token: token}
	return s.usrRepo.Create(user)
}

func (s *Service) GenerateToken(username string) (string, error) {
	const op errors.Op = "auth/service.GenerateToken"

	hash, err := bcrypt.GenerateFromPassword([]byte(username), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.E(err, errors.Internal, op)
	}

	hasher := md5.New()
	hasher.Write(hash)
	token := hex.EncodeToString(hasher.Sum(nil))
	return token, nil
}

// ActivateUser activates user
func (s *Service) ActivateUser(id uuid.UUID) error {
	const op errors.Op = "users/service.ActivateUser"

	user := &User{ID: id, Active: true}
	return s.usrRepo.Update(user)
}

// HashPassword hashes user password
func (s *Service) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// ListUsers returns all users
func (s *Service) ListUsers() ([]*User, error) {
	return s.usrRepo.Find()
}

// DeleteUser deletes a user
func (s *Service) DeleteUser(id uuid.UUID) error {
	user := &User{ID: id}
	return s.usrRepo.Delete(user)
}

func (s *Service) UpdateUser(user *User) error {
	const op errors.Op = "users/service.UpdateUser"

	return s.usrRepo.Update(user)
}
