package cache

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// RedisServer - структура, которая наследует поля объекта redis.Client и базовый контекст
type RedisServer struct {
	Client redis.Client
	Ctx    context.Context
}

// Cahe - структура, экземпляр которой является одна запись в кэше
type Cahe struct {
	Key   string
	Value int
}

// Конструктор типа RedisServer
func NewRedisServer() *RedisServer {
	redis_server := RedisServer{
		Client: *redis.NewClient(&redis.Options{Addr: "localhost:3000"}),
		Ctx:    context.Background(),
	}
	return &redis_server
}

// ChekValue выполняет поиск кэшированных данных на redis сервере
func (r *RedisServer) CheckValue(key string) (bool, error) {
	keys, err := r.Client.Exists(r.Ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return keys > 0, nil
}

// SetValue устанавливает значение по ключу в кэш redis сервера
func (r *RedisServer) SetValue(v Cahe) error {
	err := r.Client.Set(r.Ctx, v.Key, v.Value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetValue возвращает значение по ключу из кэша с redis сервера
func (r *RedisServer) GetValue(key string) (int, error) {
	value, err := r.Client.Get(r.Ctx, key).Result()
	if err != nil {
		return -1, err
	}
	sum, err := strconv.Atoi(value)
	if err != nil {
		return -1, err
	}
	return sum, nil
}
