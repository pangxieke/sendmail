package model

import (
	"strings"

	"github.com/go-redis/redis"
	"github.com/pangxieke/sendmail/config"
	"github.com/pkg/errors"
)

var (
	RedisCluster *redis.ClusterClient
	RedisClient  *redis.Client
)

func Init() (err error) {
	if err = initRedis(); err != nil {
		return
	}

	return
}

func initRedis() (err error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := RedisClient.Ping().Result()
	if err != nil {
		return errors.Wrapf(err, "Failed to connect redis")
	}
	if pong != "PONG" {
		return errors.Wrapf(err, "Failed to ping redis")
	}
	return
}

// redis cluster
func initRedisCluster() (err error) {
	opt := redis.ClusterOptions{
		Addrs:    strings.Split(config.Redis.Address, ","),
		Password: "",
		PoolSize: 10,
	}
	RedisCluster = redis.NewClusterClient(&opt)

	pong, err := RedisCluster.Ping().Result()
	if err != nil {
		return errors.Wrapf(err, "Failed to connect redis")
	}
	if pong != "PONG" {
		return errors.Wrapf(err, "Failed to ping redis")
	}
	return
}
