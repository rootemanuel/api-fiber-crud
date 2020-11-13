package cache

import (
	"log"

	"github.com/go-redis/redis"
)

var instance *GenericCache

const (
	DB_REDIS        = 0
	PASS            = ""
	URL_LOCAL_REDIS = "localhost:6379"
)

var client *redis.Client

type GenericCache struct{}

func GetClient() *redis.Client {

	if client == nil {
		nclient := redis.NewClient(&redis.Options{
			Addr:     URL_LOCAL_REDIS,
			Password: PASS,
			DB:       DB_REDIS,
		})

		_, err := nclient.Ping().Result()
		if err != nil {
			log.Println("#r00t: error open connection redis => ", err)
			return nil
		}

		client = nclient
	}

	return client
}
