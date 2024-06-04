package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `gorm:"type:TINYINT AUTO_INCREMENT;not null;primaryKey"`
	Name      string         `gorm:"type:VARCHAR(40);not null"`
	Password  string         `gorm:"type:VARCHAR(100);not null"`
	CreatedAt time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"type:DATETIME NULL;default:NULL"`
}
