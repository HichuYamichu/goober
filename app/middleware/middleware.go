package middleware

import (
	"github.com/hichuyamichu-me/uploader/internal/users"
)

type MiddlewareService struct {
	usersRepo *users.Repository
}

func NewMiddlewareService(usrRepo *users.Repository) *MiddlewareService {
	return &MiddlewareService{usrRepo}
}
