package users

import (
	"strconv"

	"github.com/hichuyamichu-me/uploader/models"
	"github.com/hichuyamichu-me/uploader/store"
)

func Register(id string, user *models.User) error {
	res, err := store.Cache.HGetAll(id).Result()
	if err != nil {
		return err
	}
	readPerm, err := strconv.ParseBool(res["read"])
	writePerm, err := strconv.ParseBool(res["write"])
	if err != nil {
		return err
	}
	user.Admin = false
	user.Write = writePerm
	user.Read = readPerm

	Create(user)
	return nil
}
