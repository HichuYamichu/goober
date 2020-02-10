package users

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// Repository performs CRUD on users table
type Repository struct {
	db *gorm.DB
}

// NewRepository creates new Repository
func NewRepository(db *gorm.DB) *Repository {
	r := &Repository{db: db}
	return r
}

// FindOne finds one user
func (r Repository) FindOne(where *User) (*User, error) {
	user := &User{}
	if r.db.Where(where).First(&user).RecordNotFound() {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Create saves a user in DB
func (r Repository) Create(user *User) error {
	if err := r.db.Create(user).Error; err != nil {
		return echo.ErrInternalServerError
	}
	return nil
}

// Delete deletes a user
func (r Repository) Delete(id int) error {
	user := &User{ID: id}
	if err := r.db.Delete(user).Error; err != nil {
		return echo.ErrInternalServerError
	}
	return nil
}

// Update updates a user
func (r Repository) Update(where *User, update *User) error {
	user := &User{}
	if err := r.db.Model(user).Where(where).Update(update).Error; err != nil {
		return echo.ErrInternalServerError
	}
	return nil
}
