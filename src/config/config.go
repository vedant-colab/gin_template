package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	} else {
		fmt.Println("Env loaded successfully")
	}

}

func GetEnv(key string, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}
