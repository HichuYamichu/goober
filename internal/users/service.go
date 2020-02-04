package users

import (
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	usrRepo *usersRepository
	cache   *redis.Client
	node    *snowflake.Node
}

func NewService(usrRepo *usersRepository, cache *redis.Client) *userService {
	node, _ := snowflake.NewNode(0)
	s := &userService{usrRepo: usrRepo, cache: cache, node: node}
	return s
}

func (s userService) CreateUser(user *User) error {
	hash, err := s.hashPassword(user.Pass)
	if err != nil {
		return err
	}
	user.Pass = hash
	s.usrRepo.Create(user)
	return nil
}

func (s userService) CreateUserFromInvite(id string, user *User) error {
	res, err := s.cache.HGetAll(id).Result()
	if err != nil {
		return err
	}
	quota, err := strconv.ParseInt(res["quota"], 10, 64)
	if err != nil {
		return err
	}
	hash, err := s.hashPassword(user.Pass)
	if err != nil {
		return err
	}
	user.Pass = hash
	user.Admin = false
	user.Quota = quota

	s.usrRepo.Create(user)
	return nil
}

func (s userService) DeleteUser(id int) {
	s.usrRepo.Delete(id)
}

func (s userService) GenereateInvite(p *UserConfig) string {
	id := s.node.Generate().String()
	s.cache.HMSet(id, "quota", p.Quota)
	s.cache.Expire(id, time.Minute*30)
	return id
}

func (s userService) VerifyCredentials(username, password string) (*User, error) {
	user := s.findOneByUsername(username)
	if user == nil {
		return nil, echo.ErrUnauthorized
	}

	match := s.checkPasswordHash(password, user.Pass)
	if !match {
		return nil, echo.ErrUnauthorized
	}

	return user, nil
}

func (s userService) findOneByUsername(username string) *User {
	user := &User{Username: username}
	user = s.usrRepo.FindOne(user)
	return user
}

func (s userService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s userService) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
