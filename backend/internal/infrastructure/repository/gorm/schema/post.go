package schema

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        int            `gorm:"type:TINYINT AUTO_INCREMENT;not null;primaryKey"`
	UserID    int            `gorm:"type:TYNYINT;not null"`
	Title     string         `gorm:"type:VARCHAR(100);not null"`
	Body      string         `gorm:"type:TEXT;not null"`
	CreatedAt time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"type:DATETIME NULL;default:NULL"`
	User      User           `gorm:"foreignKey:UserID"`
}
