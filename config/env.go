package config

import (
	"log"
	"github.com/joho/godotenv"
)

func SetupEnv() {
	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatalf("Error env %s", errEnv.Error())	
	}
}
