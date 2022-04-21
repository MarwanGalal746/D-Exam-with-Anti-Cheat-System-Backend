package driver

import (
	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
	"github.com/spf13/viper"
	"log"
)

func GetDbConnection() *rejson.Handler {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("DB_HOST") + viper.GetString("DB_PORT"),
		Password: viper.GetString("DB_PASSWORD"),
		DB:       viper.GetInt("DB_NAME"),
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Println("Can't connect with the database", "Errors:", err)
	} else {
		log.Println("Db is connected", "Errors:", err)
	}
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(client)

	return rh
}
