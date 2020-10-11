package shorty

import (
	"context"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

var (
	Gorm       *gorm.DB
	Database   *sqlx.DB
	Clickhouse *sqlx.DB
	Context    = context.Background()
)
