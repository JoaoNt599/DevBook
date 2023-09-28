package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConnMysql = ""
	ApiPort         = 0
)

func Load() {
	var er error

	if er = godotenv.Load(); er != nil {
		log.Fatal(er)
	}

	ApiPort, er = strconv.Atoi(os.Getenv("API_PORT"))
	if er != nil {
		ApiPort = 9000
	}

	StringConnMysql = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
