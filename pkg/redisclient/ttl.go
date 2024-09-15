package redisclient

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

func (c *RedisClient) TtlKey(key string) (int, error) {
	ttlDuration, err := c.con.TTL(c.ctx, key).Result()
	if err != nil && err != redis.Nil {
		return 0, errors.Wrapf(err, "ttl key %s faild", key)
	}

	if ttlDuration == -1*time.Nanosecond {
		return -1, nil
	} else {
		return int(ttlDuration.Seconds()), nil
	}
}

func (c *RedisClusterClient) TtlKey(key string) (int, error) {
	ttlDuration, err := c.con.TTL(c.ctx, key).Result()
	// fmt.Println(ttlDuration)
	// fmt.Println(int(ttlDuration.Seconds()))
	if err != nil && err != redis.Nil {
		return 0, errors.Wrapf(err, "ttl key %s faild", key)
	}
	if ttlDuration == -1*time.Nanosecond {
		return -1, nil
	} else {
		return int(ttlDuration.Seconds()), nil
	}
}
