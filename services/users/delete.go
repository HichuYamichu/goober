package users

import (
	usersRepo "github.com/hichuyamichu-me/uploader/store/users"
)

func Delete(id int) {
	usersRepo.Delete(id)
}
