package driver

import (
	"exam_service/pkg/domain/models"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func GetDbConnection() (*redis.Client, *rejson.Handler, *gorm.DB) {
	redisDb, redisJsonDb := getRedisDbConnection()
	return redisDb, redisJsonDb, getPGDbConnetion()
}

func getRedisDbConnection() (*redis.Client, *rejson.Handler) {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DB_HOST") + os.Getenv("REDIS_DB_PORT"),
		Password: os.Getenv("REDIS_DB_PASSWORD"),
		DB:       0,
	})
	_, err := redisDb.Ping().Result()
	if err != nil {
		log.Println("Can't connect with Redis database", "Errors:", err)
	} else {
		log.Println("Redis Db is connected", "Errors:", err)
	}
	redisJsonDb := rejson.NewReJSONHandler()
	redisJsonDb.SetGoRedisClient(redisDb)

	return redisDb, redisJsonDb
}

func getPGDbConnetion() *gorm.DB {
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("PG_DB_USER"),
		os.Getenv("PG_DB_PASSWORD"),
		os.Getenv("PG_DB_HOST"),
		os.Getenv("PG_DB_PORT"),
		os.Getenv("PG_DB_NAME"))
	gormDb, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Println("Can't connect with PG database", "Errors:", err)
	} else {
		log.Println("PG Db is connected", "Errors:", err)
	}
	err = gormDb.Migrator().DropTable(&models.StudentGrade{}, &models.Report{})
	if err != nil {
		panic(err)
	}
	err = gormDb.AutoMigrate(&models.StudentGrade{}, &models.Report{})
	if err != nil {
		panic(err)
	}
	return gormDb
}
