package redisclient

import "github.com/pkg/errors"

func (c *RedisClient) DelKey(key string) (int, error) {
	delCount, err := c.con.Del(c.ctx, key).Result()
	if err != nil {
		return 0, errors.Wrapf(err, "del key %s faild", key)
	}
	return int(delCount), nil
}
func (c *RedisClusterClient) DelKey(key string) (int, error) {
	delCount, err := c.con.Del(c.ctx, key).Result()
	if err != nil {
		return 0, errors.Wrapf(err, "del key %s faild", key)
	}
	return int(delCount), nil
}
