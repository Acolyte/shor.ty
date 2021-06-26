package main

import (
	"shorty/internal/app/config"
	"shorty/internal/app/shorty/api"
	"shorty/pkg/primary"
)

func main() {

	err := config.Gorm.AutoMigrate(primary.Link{})
	if err != nil {
		panic(err)
	}

	api.Serve()
}
