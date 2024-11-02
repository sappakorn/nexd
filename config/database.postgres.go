package config


import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func DatabasePostgres(cfg *Config) *gorm.DB {
	
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s TimeZone=Asia/Bangkok",
		cfg.Db.Host,			
		cfg.Db.Port,
		cfg.Db.User,
		cfg.Db.Pass,
		cfg.Db.DBName,
		
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // กำหนด output เป็น stdout
		logger.Config{
			LogLevel: logger.Info, // ระดับ log (Debug, Info, Warn, Error)
			SlowThreshold: 200 * time.Millisecond, // กำหนดเวลาที่ช้า
		},
	)

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DriverName: "pgx",
			DSN: dsn,
		}),
		&gorm.Config{
			Logger: newLogger,
		},
	)

	if err != nil {
		panic("DB connection failed" + err.Error())
	}

	return db
}
