package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(){
	if err :=godotenv.Load(); err != nil{
		log.Fatal("error while loading the env", err)
	}
}

func GetEnv(key string) string {
	 return os.Getenv(key)
}	