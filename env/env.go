package goenv

import (
	"log"
	"os"
)

// GetEnv get value form environment
func GetEnv(key string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		log.Printf("goenv.GetEnv(%s) info : key is not exist or value is empty \n", key)
	}
	return value
}
