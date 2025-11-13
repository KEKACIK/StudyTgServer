package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiHost  string
	ApiPort  int64
	ApiToken string

	BotToken string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	return &Config{
		ApiHost:  getRequiredStringEnv("API_HOST"),
		ApiPort:  getRequiredIntegerEnv("API_PORT"),
		ApiToken: getRequiredStringEnv("API_TOKEN"),

		BotToken: getRequiredStringEnv("BOT_TOKEN"),
	}
}

func getRequiredStringEnv(key string) string {
	keyValue, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Обязательная переменная %s не найдена", key)
	}
	return keyValue
}
func getRequiredIntegerEnv(key string) int64 {
	keyValue, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Обязательная переменная %s не найдена", key)
	}
	result, err := strconv.Atoi(keyValue)
	if err != nil {
		log.Fatalf("Переменная %s не является числом", key)
	}
	return int64(result)
}
