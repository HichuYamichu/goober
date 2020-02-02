package store

import (
	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

var Cache = redis.NewClient(&redis.Options{
	Addr:     viper.GetString("cache_addr"),
	Password: viper.GetString("cache_pass"),
	DB:       viper.GetInt("cache_db"),
})
