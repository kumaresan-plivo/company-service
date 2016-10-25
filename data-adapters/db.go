package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
        "github.com/hashicorp/errwrap"

	"../config"
	"fmt"
	"gopkg.in/redis.v4"
)

var db gorm.DB
var cache redis.Client

func Init() {
	c := config.GetConfig()
	connection_string := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",c.GetString("POSTGRES_DB_HOST"), c.GetString("POSTGRES_DB_USER"), c.GetString("POSTGRES_DB_NAME"), c.GetString("DB_SSL_MODE"), c.GetString("POSTGRES_DB_PASSWORD"))
	db, err := gorm.Open("postgres", connection_string)
        if errwrap.Contains(err, "credentials") {
                panic("Can't connect to database, check config!")
                // TODO handle other errors
        }
	defer db.Close()

	cache := redis.NewClient(&redis.Options{
		Addr:     c.GetString("REDIS_SERVER"),
		Password: "", // no password set
		DB:       0, // use default DB
	})

	defer cache.Close()
}

func GetDB() gorm.DB {
	return db
}

func GetCache() redis.Client {
	return cache
}
