package shorty

import (
	"gorm.io/gorm"
	"time"
)

type (
	Link struct {
		gorm.Model
		ExpiresAt time.Time
		ShortURL  string `db:"short_url" json:"shortUrl" gorm:"size:255"`
		FullURL   string `db:"full_url" json:"fullUrl" gorm:"size:255"`
	}
)
