package config

import (
	"context"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"text/template"
)

var (
	Gorm       *gorm.DB
	Database   *sqlx.DB
	Templates  map[string]*template.Template
	Clickhouse *sqlx.DB
	Context    = context.Background()
)
