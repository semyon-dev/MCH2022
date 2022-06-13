package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	MongoUrl              string
	AccessSecret          string
	DatabaseName          string
	ServerPort            string
	DefaultUserPicture    string
	DefaultProjectPicture string
	AccessTokenLifetime   int64
	RefreshTokenLifetime  int64
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("can't load from file: " + err.Error())
	}
	AccessSecret = os.Getenv("ACCESS_SECRET")
	MongoUrl = os.Getenv("MONGO_URL")
	DatabaseName = os.Getenv("DATABASE_NAME")
	ServerPort = os.Getenv("SERVER_PORT")
	RefreshTokenLifetime, err = strconv.ParseInt(os.Getenv("REFRESH_LIFETIME"), 10, 64)
	if err != nil {
		log.Fatalln("can't parse env REFRESH_LIFETIME: ", err)
	}
	AccessTokenLifetime, err = strconv.ParseInt(os.Getenv("ACCESS_LIFETIME"), 10, 64)
	if err != nil {
		log.Fatalln("can't parse env ACCESS_LIFETIME: ", err)
	}
	DefaultUserPicture = os.Getenv("DEFAULT_USER_PICTURE")
	DefaultProjectPicture = os.Getenv("DEFAULT_Project_PICTURE")
}
