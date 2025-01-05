package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg *Config) Postgres() gorm.Dialector {
	return postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Database, cfg.Password, cfg.SSLMode))
}

func NewDBConfigFromEnv() *Config {
	return &Config{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PW"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
		SSLMode:  "disable",
	}
}

func ConnRDB(cfg *Config) (*gorm.DB, error) {
	db, err := gorm.Open(cfg.Postgres())
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return db, nil
}

func NewDB() (*gorm.DB, error) {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	db.Logger = db.Logger.LogMode(logger.Info)

	if err != nil {
		return nil, err
	}
	log.Print("Connceted")
	return db, nil
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
