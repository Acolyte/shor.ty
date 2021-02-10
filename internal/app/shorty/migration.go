package shorty

import (
	"log"
	"shorty/internal/app/config"
	"shorty/pkg/shorty"
)

func Migrate() error {
	log.Println("Performing migrations")
	return config.Gorm.AutoMigrate(
		&shorty.Link{},
	)
}
