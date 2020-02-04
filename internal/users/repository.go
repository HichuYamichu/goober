package users

import (
	"github.com/jinzhu/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *usersRepository {
	r := &usersRepository{db: db}
	return r
}

func (r usersRepository) FindOne(where *User) *User {
	user := &User{}
	if err := r.db.Where(where).First(&user).Error; err != nil {
		return nil
	}
	return user
}

func (r usersRepository) Create(user *User) {
	r.db.Create(user)
}

func (r usersRepository) Delete(id int) {
	user := &User{ID: id}
	r.db.Delete(user)
}
