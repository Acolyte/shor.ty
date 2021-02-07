package shorty

import (
	"shorty/internal/app/config"
	"shorty/pkg/shorty"
)

func Migrate() error {
	return config.Gorm.AutoMigrate(
		&shorty.Link{},
	)
}
