package cache

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/rootemanuel/api-fiber-crud/entity"
)

type AccountCache struct {
	GenericCache
}

func (m *AccountCache) Teste() *entity.AccountEntity {
	client := GetClient()

	s, err := client.Get("root").Result()
	if err == redis.Nil {
		fmt.Println("root does not exist")
	}

	usr := &entity.AccountEntity{}
	err = json.Unmarshal([]byte(s), &usr)
	if err != nil {
		log.Println(err)
	}

	return usr
}
