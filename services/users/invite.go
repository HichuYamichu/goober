package users

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/hichuyamichu-me/uploader/store"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(0)
	if err != nil {
		panic(err)
	}
}

type UserConfig struct {
	Quota int64 `json:"quota"`
}

func GenereateInvite(p *UserConfig) string {
	id := node.Generate().String()
	store.Cache.HMSet(id, "quota", p.Quota)
	store.Cache.Expire(id, time.Minute*30)
	return id
}
