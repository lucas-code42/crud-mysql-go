package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var StringConnection = ""

func LoadEnviroment() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	StringConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB"),
	)
	fmt.Println(StringConnection)
}
