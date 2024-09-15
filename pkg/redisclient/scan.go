package redisclient

import (
	"bufio"
	"context"

	// "fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type ScanKeys interface {
	ScanKeys(string) ([]string, error)
	ScanKeysToFile(searchKey, fPath string) (int, error)
}

// scan redis all keys to return
func (c *RedisClient) ScanKeys(searchKey string) ([]string, error) {
	var allKeys []string
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = c.con.Scan(c.ctx, cursor, searchKey, 100).Result()
		if err != nil {
			return nil, errors.Wrapf(err, "scan %s failed", searchKey)
		}
		if cursor == 0 {
			break
		}
		allKeys = append(allKeys, keys...)
	}
	return allKeys, nil
}

func (c *RedisClusterClient) ScanKeys(searchKey string) ([]string, error) {
	var allKeys []string
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = c.con.Scan(c.ctx, cursor, searchKey, 100).Result()
		if err != nil {
			return nil, errors.Wrapf(err, "scan %s failed", searchKey)
		}
		if cursor == 0 {
			break
		}
		allKeys = append(allKeys, keys...)
	}
	return allKeys, nil
}

// scan redis all keys output to file
func (c *RedisClient) ScanKeysToFile(searchKey, fPath string) (int, error) {
	count := 0
	f, err := os.OpenFile(fPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	buf := bufio.NewWriter(f)

	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = c.con.Scan(c.ctx, cursor, searchKey, 100).Result()
		if err != nil {
			return 0, errors.Wrapf(err, "scan %s failed", searchKey)
		}
		for _, key := range keys {
			count++
			buf.WriteString(key + "\n")
		}
		if cursor == 0 {
			break
		}
	}
	return count, buf.Flush()
}

func (c *RedisClusterClient) ScanKeysToFile(searchKey, fPath string) (int, error) {
	count := 0
	f, err := os.OpenFile(fPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	buf := bufio.NewWriter(f)

	err = c.con.ForEachMaster(c.ctx, func(ctx context.Context, con *redis.Client) error {
		var cursor uint64
		for {
			var keys []string
			var err error
			keys, cursor, err = con.Scan(ctx, cursor, searchKey, 100).Result()
			if err != nil {
				return errors.Wrapf(err, "scan %s failed", searchKey)
			}
			for _, key := range keys {
				count++
				buf.WriteString(key + "\n")
			}
			if cursor == 0 {
				break
			}
		}
		return buf.Flush()
	})

	if err != nil {
		return 0, err
	}
	// fmt.Println(count)
	return count, buf.Flush()

}

// func (c *RedisClusterClient) ScanKeysToFile(searchKey, fPath string) (int, error) {
// 	count := 0
// 	f, err := os.OpenFile(fPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer f.Close()

// 	buf := bufio.NewWriter(f)

// 	// slots, err := c.con.ClusterSlots(c.ctx).Result()
// 	// if err != nil {
// 	// return 0, errors.Wrapf(err, "get all nodes slots faild")
// 	// }

// 	var cursor uint64
// 	for {
// 		var keys []string
// 		var err error
// 		keys, cursor, err = c.con.Scan(c.ctx, cursor, searchKey, 100).Result()
// 		if err != nil {
// 			return 0, errors.Wrapf(err, "scan %s failed", searchKey)
// 		}
// 		for _, key := range keys {
// 			count++
// 			buf.WriteString(key + "\n")
// 		}
// 		if cursor == 0 {
// 			break
// 		}
// 	}
// 	return count, buf.Flush()
// }
