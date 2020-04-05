package middleware

import (
	"github.com/hichuyamichu-me/goober/domain/users"
)

type Service struct {
	usersRepo *users.Repository
}

func NewService(usrRepo *users.Repository) *Service {
	return &Service{usrRepo}
}
