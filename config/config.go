package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() {
	err:=godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GetString(key string) string {
	value,ok := os.LookupEnv(key)
	if !ok {
		panic("key not found")
	}
	return value
}

func GetInt(key string) int {
	value := GetString(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return i
}




