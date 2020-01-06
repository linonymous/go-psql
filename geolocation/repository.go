package geolocation

import "github.com/jinzhu/gorm"

type GeolocationRepository struct {
	database *gorm.DB
}

func NewGeolocationRepository(db *gorm.DB) *GeolocationRepository {
	return &GeolocationRepository{database: db}
}
func (repo *GeolocationRepository) Query() {
	// code to query
	return
}
