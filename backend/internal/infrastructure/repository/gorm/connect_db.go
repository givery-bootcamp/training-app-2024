package gorm

import (
	"context"
	"database/sql"
	"fmt"

	"myapp/internal/config"
	"myapp/internal/infrastructure/repository/gorm/migrate"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbgen *DB

type DB struct {
	db *gorm.DB
}

func GetQuery(ctx context.Context) *gorm.DB {
	txq := getQueryWithContext(ctx)
	if txq != nil {
		return txq
	}
	return dbgen.db
}

func Set(db *gorm.DB) *DB {
	return &DB{db: db}
}

func (d *DB) Transaction(ctx context.Context, txOps *sql.TxOptions, fn func(ctx context.Context) error) error {
	fc := func(tx *gorm.DB) error {
		ctxWithQuery := withQuery(ctx, tx)

		err := fn(ctxWithQuery)
		if err != nil {
			return err
		}

		return nil
	}

	if txOps == nil {
		err := d.db.Transaction(fc)
		if err != nil {
			return err
		}
	} else {
		err := d.db.Transaction(fc, txOps)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewGormDB(ctx context.Context, config config.DBConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2FTokyo&charset=utf8mb4",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")
	if err != nil {
		panic(err)
	}
	err = migrate.Migrate(ctx, db)
	if err != nil {
		panic(fmt.Errorf("failed to migrate: %w", err))
	}

	dbgen = Set(db)
}

type CtxKey string

const (
	QueryKey CtxKey = "query"
)

func withQuery(ctx context.Context, query *gorm.DB) context.Context {
	return context.WithValue(ctx, QueryKey, query)
}

func getQueryWithContext(ctx context.Context) *gorm.DB {
	query, ok := ctx.Value(QueryKey).(*gorm.DB)
	if !ok {
		return nil
	}
	return query.Session(&gorm.Session{
		Context: ctx,
	})
}
