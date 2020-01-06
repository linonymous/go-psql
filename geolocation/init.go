package geolocation

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func ConnectDatabase(config *Config) (*gorm.DB, error) {
	gormConf := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", config.Host, config.Port, config.User, config.DBName, config.SSLMode, config.Password)
	return gorm.Open(config.DBDriver, gormConf)
}
