package database

import (
	"day6/pkg/util"
	"sync"

	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

func CreateConnection() {
	conf := dbConfig{
		User: util.GetEnv("DB_USER", "root"),
		Pass: util.GetEnv("DB_PASS", ""),
		Host: util.GetEnv("DB_HOST", "localhost"),
		Port: util.GetEnv("DB_PORT", "3306"),
		Name: util.GetEnv("DB_NAME", "training"),
	}

	mysql := mysqlConfig{dbConfig: conf}
	once.Do(func() {
		mysql.Connect()
	})
}

func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}
func InitMigrate() {
	dbConn.AutoMigrate()
}
