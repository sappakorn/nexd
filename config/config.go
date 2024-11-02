package config

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/gorm"
)

type App struct {
	Port int
}

type Db struct {
	Host 		string 
	Port 		int
	User 		string
	Pass 		string 
	DBName 		string 
	SSLMode 	string
	TimeZone 	string
}

type Config struct {
	App App
	Db Db
}

type Database interface {
	GetDb() *gorm.DB
}

func GetConfig() Config {
	dbPort, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	// appPort, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	dbInformation := Db{}
	fmt.Print("APP_ENV : ", os.Getenv("APP_ENV"))
	if os.Getenv("APP_ENV") == "development" { 
		dbInformation.Host = os.Getenv("DB_HOST")
		dbInformation.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
		dbInformation.User = os.Getenv("DB_USER")
		dbInformation.Pass = os.Getenv("DB_PASS")
		dbInformation.DBName = os.Getenv("DB_NAME")
		dbInformation.SSLMode = os.Getenv("DB_SSLMODE")
		dbInformation.TimeZone = os.Getenv("DB_TIMEZONE")
	}else{
		dbInformation.Host = os.Getenv("DB_HOST")
		dbInformation.Port,_ = strconv.Atoi(os.Getenv("DB_PORT"))
		dbInformation.User = os.Getenv("DB_USER")
		dbInformation.Pass = os.Getenv("DB_PASS")
		dbInformation.DBName = os.Getenv("DB_NAME")
		dbInformation.SSLMode = os.Getenv("DB_SSLMODE")
		dbInformation.TimeZone = os.Getenv("DB_TIMEZONE")
	}

	return Config{
		App: App{
			Port: dbPort,
		},
		Db: dbInformation,
	}
}