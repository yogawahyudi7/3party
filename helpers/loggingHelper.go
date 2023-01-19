package helpers

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"

	Config "3party/config"
)

func LoggingCloudPubSub(wg *sync.WaitGroup, params string, channelLog string) (status int, response string) {
	defer wg.Done()
	var ctx = context.Background()

	redisUrl, _ := Config.RedisUrl()
	redisPassword, _ := Config.RedisPassword()
	redisDB, _ := Config.RedisDB()

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPassword, // no password set
		DB:       redisDB,       // use default DB
	})

	err := rdb.Publish(ctx, channelLog, params).Err()
	if err != nil {
		return 400, err.Error()
	}

	return 200, "Success Publish Logging"
}
