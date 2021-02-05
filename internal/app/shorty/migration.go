package shorty

import (
	"shorty/pkg/shorty"
)

func Migrate() error {
	return Gorm.AutoMigrate(
		&shorty.Link{},
	)
}
