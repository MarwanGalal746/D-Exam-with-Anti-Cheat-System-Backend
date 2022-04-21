package driver

import (
	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
	"github.com/spf13/viper"
	"log"
)

func GetDbConnection() (*redis.Client, *rejson.Handler) {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("DB_HOST") + viper.GetString("DB_PORT"),
		Password: viper.GetString("DB_PASSWORD"),
		DB:       viper.GetInt("DB_NAME"),
	})
	_, err := redisDb.Ping().Result()
	if err != nil {
		log.Println("Can't connect with the database", "Errors:", err)
	} else {
		log.Println("Db is connected", "Errors:", err)
	}
	redisJsonDb := rejson.NewReJSONHandler()
	redisJsonDb.SetGoRedisClient(redisDb)

	return redisDb, redisJsonDb
}
