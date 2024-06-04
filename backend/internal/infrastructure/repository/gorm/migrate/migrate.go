package migrate

import (
	"context"
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migrations = []*gormigrate.Migration{
	v1(),
}

func Migrate(ctx context.Context, db *gorm.DB) error {
	m := gormigrate.New(db.Session(&gorm.Session{Context: ctx}), gormigrate.DefaultOptions, migrations)

	if err := m.Migrate(); err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}
	return nil
}
