package config

import (
	"log"
	"os"
)

func GetEnv(envName string) string {
	envValue, envExists := os.LookupEnv(envName)
	if !envExists {
		log.Fatalf("error getting env: %s", envName)
	}
	return envValue
	
}