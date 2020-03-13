package auth

import (
	"github.com/hichuyamichu-me/uploader/errors"
	"github.com/hichuyamichu-me/uploader/internal/users"
	"golang.org/x/crypto/bcrypt"
)

// Service performs operations specyfic to user domain
type Service struct {
	usrRepo *users.Repository
}

// NewService creates new user service
func NewService(usrRepo *users.Repository) *Service {
	s := &Service{usrRepo: usrRepo}
	return s
}

// VerifyCredentials verifies user credentials
func (s Service) VerifyCredentials(username, password string) (*users.User, error) {
	const op errors.Op = "auth/service.VerifyCredentials"

	user, err := s.usrRepo.FindByUsername(username)
	if err != nil {
		return nil, errors.E(err, errors.Authentication, op)
	}

	match := s.checkPasswordHash(password, user.Pass)
	if !match {
		return nil, errors.E(errors.Errorf("invalid password"), errors.Authentication, op)
	}

	return user, nil
}

func (s Service) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
