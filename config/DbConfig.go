package config

import (
	"blogapi/internal/models"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	DbHost   string
	DbPort   string
	UserName string
	Password string
	DbName   string
}

func LoadConfig() (*DbConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	return &DbConfig{
		DbHost:   GetEnvOrDefault("DB_HOST", "127.0.0.1"),
		DbPort:   GetEnvOrDefault("DB_PORT", "3306"),
		DbName:   GetEnvOrDefault("DB_NAME", "blogapi"),
		UserName: GetEnvOrDefault("DB_USER_NAME", "root"),
		Password: GetEnvOrDefault("DB_PASSWORD", ""),
	}, nil
}

func ConnectWithDb() *gorm.DB {
	dbConfig, err := LoadConfig()
	if err != nil {
		panic("Failed to load config")
	}
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.DbHost,
		dbConfig.DbPort,
		dbConfig.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Falied to create connection db")
	}
	db.AutoMigrate(&models.Comment{}, &models.Post{}, &models.User{})
	return db
}

func CloseDbConnection(db *gorm.DB){
	dbSql, err:=db.DB()
	if err!=nil{
		panic("Failed to close connection db")
	}
	dbSql.Close()
}
