package users

import (
	"github.com/hichuyamichu-me/uploader/errors"
	"github.com/jinzhu/gorm"
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

// FindOne finds one user
func (r Repository) FindOne(where *User) (*User, error) {
	const op errors.Op = "users/repository.FindOne"

	user := &User{}
	if r.db.Where(where).First(&user).RecordNotFound() {
		return nil, errors.E(errors.Errorf("user not found"), errors.NotFound, op)
	}
	return user, nil
}

// Create saves a user in DB
func (r Repository) Create(user *User) error {
	const op errors.Op = "users/repository.Create"

	if err := r.db.Create(user).Error; err != nil {
		return errors.E(err, errors.Internal, op)
	}
	return nil
}

// Delete deletes a user
func (r Repository) Delete(id int) error {
	const op errors.Op = "users/repository.Delete"

	user := &User{ID: id}
	if err := r.db.Delete(user).Error; err != nil {
		return errors.E(err, errors.Internal, op)
	}
	return nil
}

// Update updates a user
func (r Repository) Update(u *User) error {
	const op errors.Op = "users/repository.Update"

	// user := &User{}
	if err := r.db.Model(u).Update(u).Error; err != nil {
		return errors.E(err, errors.Internal, op)
	}
	return nil
}
