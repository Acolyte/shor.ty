package user

import (
	"gorm.io/gorm"
	"time"
)

type (
	User struct {
		ID        uint           `gorm:"primarykey"`
		Name      string         `db:"name"`
		Username  string         `db:"username"`
		Password  string         `db:"password"`
		CreatedAt time.Time      `db:"created_at"`
		UpdatedAt time.Time      `db:"updated_at"`
		DeletedAt gorm.DeletedAt `gorm:"index" db:"deleted_at"`
	}
)
