package redisclient

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type RedisConnectSettingS struct {
	Host         string
	Port         string
	Password     string
	DB           int
	PoolSize     int
	MinIdleConns int
}

type RedisClient struct {
	con *redis.Client
	// con *redis.Client|*redis.ClusterClient

	// ctx context.Context
	// con interface{}
	ctx context.Context
}

type RedisClusterClient struct {
	con *redis.ClusterClient
	ctx context.Context
}

func NewRedisConnect(host, port string, password string, db int, poolSize int, minIdleConnect int) (*RedisClient, error) {
	ctx := context.Background()
	addr := fmt.Sprintf("%s:%s", host, port)
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           db,
		PoolSize:     poolSize,
		MinIdleConns: minIdleConnect,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "connect redis %s faile", addr)
	}

	return &RedisClient{
		client, ctx,
	}, nil
}

func NewRedisClusterConnect(host, port string, password string, db int, poolSize int, minIdleConnect int) (*RedisClusterClient, error) {
	ctx := context.Background()
	addr := fmt.Sprintf("%s:%s", host, port)
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        []string{addr},
		Password:     password,
		PoolSize:     poolSize,
		MinIdleConns: minIdleConnect,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "connect redis %s faile", addr)
	}

	return &RedisClusterClient{
		client, ctx,
	}, nil
	// return &RedisClient{
	// client, ctx,
	// }, nil
}
