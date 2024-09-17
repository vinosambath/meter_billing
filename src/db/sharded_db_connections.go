package db

import (
	"MeterBilling/src/configuration"
	"MeterBilling/src/logger"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConnections []*gorm.DB
var dbConnectionsSingleton sync.Once

type ShardedDBConnections struct {
}

func InitDBConnections() {
	dbConnectionsSingleton.Do(func() {
		config := configuration.GetConfig()
		for _, database := range config.Database.Connections {
			dbConn, err := gorm.Open(mysql.Open(database.DSN), &gorm.Config{})
			if err != nil {
				logger.GetLogger().Fatal("error initialising db connection")
				panic(err)
			}
			dbConnections = append(dbConnections, dbConn)
		}
	})
}

func GetDBConnections() []*gorm.DB {
	return dbConnections
}
