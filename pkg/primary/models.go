package primary

import (
	"gorm.io/gorm"
	"time"
)

type (
	Link struct {
		ID        uint   `gorm:"primarykey"`
		UUID      string `db:"uuid" gorm:"size:32, type:varchar, uniqueIndex"`
		FullURL   string `db:"full_url" gorm:"size:2048"`
		Scheme    string `db:"scheme" gorm:"size:8,index:idx_scheme"`
		Host      string `db:"host" gorm:"size:255, index:idx_host"`
		Port      int    `db:"host" gorm:"index:idx_port"`
		Path      string `db:"path" gorm:"index:idx_path"`
		Query     string `db:"query" gorm:"index:idx_query"`
		ExpiresIn string `db:"-"`

		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
		ExpiresAt time.Time      `db:"expires_at" json:"expiresAt"`
	}

	FoundViewData struct {
		Link
		HostURL string
	}

	ErrorViewData struct {
		Message string
	}
)
