package data_access

import (
	"context"
	"server/common"
	"server/model"

	"github.com/redis/go-redis/v9"
)

const (
	usersPrefix  = "user:"
	usersListKey = "users"
	usernameKey  = "username"
	passwordKey  = "password"
	userIdKey    = "userId"
)

type UserStore struct {
	client *redis.Client
}

func NewUserStore() *UserStore {
	addr := common.GetEnv(common.RedisUrl)
	password := common.GetEnv(common.RedisPassword)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return &UserStore{client: client}
}

func (s *UserStore) Insert(user *model.User) error {
	ctx := context.Background()

	err := s.client.HSet(ctx, usersPrefix+user.ID, usernameKey, user.Username, passwordKey, user.Password).Err()
	if err != nil {
		return err
	}

	err = s.client.LPush(ctx, usersListKey, user.ID).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) GetByID(id string) (*model.User, error) {
	ctx := context.Background()

	result, err := s.client.HGetAll(ctx, usersPrefix+id).Result()
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:       id,
		Username: result[usernameKey],
		Password: result[passwordKey],
	}

	return user, nil
}

func (s *UserStore) GetByUsername(username string) (*model.User, error) {
	ctx := context.Background()

	userIds, err := s.client.LRange(ctx, usersListKey, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	for _, userId := range userIds {
		user, err := s.GetByID(userId)
		if err != nil {
			return nil, err
		}

		if user.Username == username {
			return user, nil
		}
	}

	return nil, nil
}
