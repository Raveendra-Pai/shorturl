package main

import (
	"errors"
	"strconv"

	"github.com/go-redis/redis"
)

type RedisMap struct {
	redisClient *redis.Client
}

func (r *RedisMap) Init(conf Config) error {

	redisurl := conf.Redis.Ip + ":" + strconv.FormatUint(uint64(conf.Redis.Port), 10)

	Applog.Info("Forming Redis connect url : " + redisurl)

	r.redisClient = redis.NewClient(&redis.Options{
		Addr: redisurl,
	})

	_, err := r.redisClient.Ping().Result()

	if err != nil {
		Applog.Error("Unable to send Ping Redis server: " + redisurl)
		return err
	}
	Applog.Info("Succesfully connected to redis server: " + redisurl)
	return nil
}

func (r *RedisMap) Insert(key string, value string) error {

	var err error = nil
	defer func() {
		if er := recover(); er != nil {
			Applog.Error("panic occurred in Redis Insert")
			err = errors.New("failed to insert into redis for key: " + key + " value: " + value)
		}
	}()
	r.redisClient.Set(key, value, 0)
	return err
}

func (r *RedisMap) Retrieve(key string) (string, error) {
	var er error = nil
	defer func() {
		if er := recover(); er != nil {
			Applog.Error("panic occurred in Redis Get")
			er = errors.New("failed to insert into redis for key: " + key)
		}
	}()

	longurl, err := r.redisClient.Get(key).Result()

	if err != nil {
		return "", err
	}
	er = err
	return longurl, er
}
