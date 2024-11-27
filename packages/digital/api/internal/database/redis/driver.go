package redis

import (
	"context"
	"fmt"
	"github.com/cyberpunkattack/environment"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type RedisDB struct {
	instance *redis.Client
	addr string
	password string
	db int
}

var redisDbInstance *RedisDB = nil



func New() *RedisDB {
	REDIS_HOST := environment.GEnv().GetVariable("REDIS_URI")
	REDIS_PASSWORD := environment.GEnv().GetVariable("REDIS_PASSWORD")
	REDIS_DBNAME, _ := strconv.Atoi(environment.GEnv().GetVariable("REDIS_DBNAME"))
	rdb := redis.NewClient(&redis.Options{
        Addr:     REDIS_HOST,
        Password: REDIS_PASSWORD, // no password set
        DB:       REDIS_DBNAME,  // use default DB
    })

	redisDbInstance = &RedisDB{
		instance: rdb,
		addr: REDIS_HOST,
		password: REDIS_PASSWORD,
		db: REDIS_DBNAME,
	}

	redisDbInstance.Get().Set(context.Background(), "hello", "world", 5000)
	res := redisDbInstance.Get().Get(context.Background(), "hello").String()
	fmt.Println(res);


	return redisDbInstance
}


func Db() *RedisDB {
	return redisDbInstance
}

func (rd* RedisDB) Get() *redis.Client {
	return rd.instance
}