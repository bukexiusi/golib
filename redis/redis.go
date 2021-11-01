package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type (
	Redis struct {
		cli          *redis.Client
		luaScriptMap LuaScriptMap
	}

	Configuration struct {
		Addr     string
		Password string
		DB       int
	}

	LuaScriptMap map[string]string
)

func Dial(config *Configuration) (*Redis, error) {
	result := new(Redis)

	// New client
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, errors.Wrap(err, "redis - Dial - Ping")
	}
	result.cli = rdb

	// Init lua scripts
	initLuaScriptMap(result)

	return result, nil
}

func (r *Redis) Close() {
	_ = r.cli.Close()
}

func (r *Redis) GetDel(key string) (string, error) {
	result, err := r.cli.EvalSha(r.cli.Context(), r.luaScriptMap["getdel"], []string{key}).Result()
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

func (r *Redis) HsetEx(key string, argv ...string) (int64, error) {
	result, err := r.cli.EvalSha(r.cli.Context(), r.luaScriptMap["hsetex"], []string{key}, argv).Result()
	if err != nil {
		return -1, err
	}
	return result.(int64), nil
}

func (r *Redis) Get(key string) (string, error) {
	result, err := r.cli.Get(r.cli.Context(), key).Result()
	if err == redis.Nil {
		err = nil
	}
	return result, err
}

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	return r.cli.Set(r.cli.Context(), key, value, expiration).Err()
}

func (r *Redis) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return r.cli.SetNX(r.cli.Context(), key, value, expiration).Result()
}

func (r *Redis) Del(key ...string) (int64, error) {
	return r.cli.Del(r.cli.Context(), key...).Result()
}

func (r *Redis) HGetAll(key string) (map[string]string, error) {
	return r.cli.HGetAll(r.cli.Context(), key).Result()
}
