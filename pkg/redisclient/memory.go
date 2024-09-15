package redisclient

import (
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

// reports the number of bytes that a key and its value require to be stored
func (c *RedisClient) MemoryUsage(key string, samples int) (int, error) {
	size, err := c.con.MemoryUsage(c.ctx, key, samples).Result()
	if err != nil && err != redis.Nil {
		return 0, errors.Wrapf(err, "memory key %s faild", key)
	}

	return int(size), nil
}
func (c *RedisClusterClient) MemoryUsage(key string, samples int) (int, error) {
	size, err := c.con.MemoryUsage(c.ctx, key, samples).Result()
	if err != nil && err != redis.Nil {
		return 0, errors.Wrapf(err, "memory key %s faild", key)
	}

	return int(size), nil
}
