package users

import (
	"fmt"

	"github.com/hichuyamichu-me/uploader/errors"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
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

// FindByUsername finds user by username
func (r Repository) FindByUsername(username string) (*User, error) {
	const op errors.Op = "users/repository.FindByUsername"

	user := &User{}
	if err := r.db.Where(&User{Username: username}).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.E(errors.Errorf("user not found"), errors.NotFound, op)
		}
		return nil, errors.E(err, errors.Internal, op)
	}

	return user, nil
}

// FindByToken finds user by username
func (r Repository) FindByToken(token string) (*User, error) {
	const op errors.Op = "users/repository.FindByToken"

	user := &User{}
	if err := r.db.Where(&User{Token: token}).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.E(errors.Errorf("user not found"), errors.NotFound, op)
		}
		return nil, errors.E(err, errors.Internal, op)
	}

	return user, nil
}

// Find returns all users
func (r Repository) Find() ([]*User, error) {
	const op errors.Op = "users/repository.Find"

	users := []*User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, errors.E(err, errors.Internal, op)
	}

	return users, nil
}

// Create saves a user in DB
func (r Repository) Create(u *User) error {
	const op errors.Op = "users/repository.Create"

	if err := r.db.Create(u).Error; err != nil {
		e, ok := err.(*pq.Error)
		if ok {
			fmt.Println(e.Code)
			if e.Code == "23505" {
				return errors.E(err, errors.Invalid, op)
			}
		}
		return errors.E(err, errors.Internal, op)
	}

	return nil
}

// Update updates a user
func (r Repository) Update(u *User) error {
	const op errors.Op = "users/repository.Update"

	if err := r.db.Model(u).Update(u).Error; err != nil {
		return errors.E(err, errors.Internal, op)
	}

	return nil
}

// Delete deletes a user
func (r Repository) Delete(u *User) error {
	const op errors.Op = "users/repository.Delete"

	if err := r.db.Delete(u).Error; err != nil {
		return errors.E(err, errors.Internal, op)
	}

	return nil
}
