package primary

import (
	"database/sql"
	"time"
)

type (
	Link struct {
		ID   uint   `gorm:"primarykey" json:"-"`
		UUID string `db:"uuid" json:"uuid" gorm:"size:32, uniqueIndex"`

		CreatedAt time.Time    `db:"created_at" json:"-"`
		UpdatedAt time.Time    `db:"updated_at" json:"-"`
		DeletedAt sql.NullTime `db:"deleted_at" gorm:"index" json:"-"`
		ExpiresIn string       `db:"-" json:"expireIn"`
		ExpiresAt time.Time    `db:"expires_at" json:"expiresAt"`

		FullURL string `db:"full_url" json:"fullUrl" gorm:"size:2048"`
		Scheme  string `db:"scheme" json:"-" gorm:"size:8,index:idx_scheme"`
		Host    string `db:"host" json:"-" gorm:"size:255, index:idx_host"`
		Port    int    `db:"host" json:"-" gorm:"index:idx_port"`
		Path    string `db:"path" json:"-" gorm:"index:idx_path"`
		Query   string `db:"query" json:"-" gorm:"index:idx_query"`
	}

	FoundViewData struct {
		Link
		HostURL string
	}
)
