package dao

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() {
	err := connectMySQL()
	if err != nil {
		panic(err)
	}
	err = connectRedis()
	if err != nil {
		panic(err)
	}
}

var (
	DB  *gorm.DB
	RDB *redis.Client
)

var Ctx = context.Background()

func connectMySQL() error {
	dsn := "root:p@ssw0rd@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func connectRedis() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		return err
	}
	RDB = rdb
	return nil
}
