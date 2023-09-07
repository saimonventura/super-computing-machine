package infrastructure

import (
	"context"
	"encoding/json"
	"super-computing-machine/auth/domain"

	"github.com/go-redis/redis/v8"
)

type RedisUserRepository struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisUserRepository(c *redis.Client, ctx context.Context) *RedisUserRepository {
	return &RedisUserRepository{client: c, ctx: ctx}
}

func (r *RedisUserRepository) FindByEmail(email string) (*domain.User, error) {
	data, err := r.client.Get(r.ctx, email).Result()
	if err != nil {
		return nil, err
	}
	var user domain.User
	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RedisUserRepository) Save(user *domain.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.client.Set(r.ctx, user.Email, data, 0).Err()
}
