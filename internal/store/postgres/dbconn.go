package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDBDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Europe/Moscow",
		viper.Get("db.host"),
		viper.Get("db.user"),
		viper.Get("db.password"),
		viper.Get("db.name"),
		viper.Get("db.port"),
	)
}

type moneyServiceRepo struct {
	db *gorm.DB
}

func NewMoneyServiceRepo() *moneyServiceRepo {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(getDBDSN()), &gorm.Config{Logger: newLogger})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &moneyServiceRepo{db: db}
}
