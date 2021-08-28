package support

import (
	"log"
	"os"
	"strconv"
)

func GetEnvStr(env string, def string) string {
	value := os.Getenv(env)
	if value == "" {
		return def
	}
	return value
}

func GetEnvInt(env string, def int) int {
	value, err := strconv.Atoi(os.Getenv(env))
	if err != nil {
		log.Print("Couldn't parse")
		return def
	}
	return value
}
