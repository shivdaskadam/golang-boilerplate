package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	DBName   string
	Port     string
}

func InitDatabase() (*gorm.DB, error) {
	Username := Instance().SQLUsername
	Password := Instance().SQLPassword
	Host := Instance().SQLHost
	DBName := Instance().SQLDBName
	Port := Instance().SQLPort

	dsn := Username + ":" + Password + "@tcp(" + Host + ":" + Port + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Open the database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
