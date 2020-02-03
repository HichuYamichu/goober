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
	quota, err := strconv.ParseInt(res["quota"], 10, 64)
	if err != nil {
		return err
	}
	user.Admin = false
	user.Quota = quota

	Create(user)
	return nil
}
