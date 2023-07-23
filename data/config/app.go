package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var _database *gorm.DB

func Connect() {
	connString := "Server=localhost;Database=API;User Id=SA;Password=yourStrong(!)Password"
	database, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	_database = database
}

func GetDatabase() *gorm.DB {
	return _database
}
