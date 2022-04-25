package driver

import (
	"exam_service/pkg/domain/models"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func GetDbConnection() (*redis.Client, *rejson.Handler, *gorm.DB) {
	redisDb, redisJsonDb := getRedisDbConnection()
	return redisDb, redisJsonDb, getPGDbConnetion()
}

func getRedisDbConnection() (*redis.Client, *rejson.Handler) {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_DB_HOST") + viper.GetString("REDIS_DB_PORT"),
		Password: viper.GetString("REDIS_DB_PASSWORD"),
		DB:       viper.GetInt("REDIS_DB_NAME"),
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
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("PG_DB_HOST"),
		viper.GetString("PG_DB_PORT"),
		viper.GetString("PG_DB_USER"),
		viper.GetString("PG_DB_PASSWORD"),
		viper.GetString("PG_DB_NAME"),
	)
	gormDb, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
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
