package main

import (
	"github.com/go-psql/geolocation"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func main() {
	config := geolocation.NewConfig()
	db, err := geolocation.ConnectDatabase(config)
	if err != nil {
		log.Fatalf("Couldn't start the database %s", err.Error())
	}

	repository := geolocation.NewGeolocationRepository(db)
	service := geolocation.NewGeolocationService(config, repository)

	service.Query()
}
