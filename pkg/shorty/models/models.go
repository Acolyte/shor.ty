package models

import (
	"gorm.io/gorm"
)

type (
	Link struct {
		gorm.Model
		ShortURL string
		FullURL  string
	}
)
