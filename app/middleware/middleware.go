package middleware

import (
	"github.com/hichuyamichu-me/goober/internal/users"
)

type MiddlewareService struct {
	usersRepo *users.Repository
}

func NewMiddlewareService(usrRepo *users.Repository) *MiddlewareService {
	return &MiddlewareService{usrRepo}
}
