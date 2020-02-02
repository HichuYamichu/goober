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

type Permissions struct {
	Write bool `json:"write"`
	Read  bool `json:"read"`
}

func GenereateInvite(p *Permissions) string {
	id := node.Generate().String()
	store.Cache.HMSet(id, "write", p.Write, "read", p.Read)
	store.Cache.Expire(id, time.Minute*30)
	return id
}
