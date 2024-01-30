package common

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	_ "github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func GetDbConnection() *gorm.DB {
	//dbHost := os.Getenv("DB_HOST")
	//dbUser := os.Getenv("DB_USER")
	//dbPassword := os.Getenv("DB_PASSWORD")
	//dbPort := os.Getenv("DB_PORT")
	//dbName := os.Getenv("DB_NAME")

	url := fmt.Sprintf("host=localhost user=admin password=password123 dbname=uwang-dev port=5432 sslmode=disable TimeZone=Asia/Jakarta")

	database, _ := gorm.Open(postgres.Open(url), &gorm.Config{})
	return database
}

func GetRedisConnection() *redis.Client {
	redisAdd := os.Getenv("REDIS_ADDRESS")
	redisUser := os.Getenv("REDIS_USER")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, _ := strconv.ParseInt(os.Getenv("REDIS_DB"), 0, 8)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAdd,
		Username: redisUser,
		Password: redisPassword,
		DB:       int(redisDB),
	})

	return redisClient
}
