package driver

import (
	"github.com/go-redis/redis"
	"log"
)

func GetDbConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Println("Can't connect with the database", "Errors:", err)
	} else {
		log.Println("Db is connected", "Errors:", err)
	}
	return client
}
