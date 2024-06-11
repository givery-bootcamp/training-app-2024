package migrate

import (
	"myapp/internal/infrastructure/repository/gorm/schema"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func v1() *gormigrate.Migration {
	tables := []interface{}{
		&schema.HelloWorld{},
		&schema.User{},
		&schema.Post{},
	}
	return &gormigrate.Migration{
		ID: "1",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.AutoMigrate(tables...); err != nil {
				return err
			}

			if err := firstInsertHelloWorld(tx); err != nil {
				return err
			}

			if err := firstInsertUser(tx); err != nil {
				return err
			}

			if err := firstInsertPost(tx); err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(tables...)
		},
	}
}

func firstInsertHelloWorld(db *gorm.DB) error {
	helloWorlds := []schema.HelloWorld{
		{
			Lang:    "en",
			Message: "Hello, World!",
		},
		{
			Lang:    "ja",
			Message: "こんにちは、世界！",
		},
	}

	for _, hello := range helloWorlds {
		err := db.
			Where("lang = ?", hello.Lang).
			FirstOrCreate(&hello).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func firstInsertUser(db *gorm.DB) error {
	users := []schema.User{
		{
			Name:     "taro",
			Password: "password",
		},
		{
			Name:     "hanako",
			Password: "password",
		},
	}

	for _, user := range users {
		err := db.
			Where("name = ?", user.Name).
			FirstOrCreate(&user).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func firstInsertPost(db *gorm.DB) error {
	posts := []schema.Post{
		{
			UserID: 1,
			Title:  "test1",
			Body:   "質問1\n改行",
		},
		{
			UserID: 1,
			Title:  "test2",
			Body:   "質問2\n改行",
		},
	}

	for _, post := range posts {
		err := db.
			Where("title = ?", post.Title).
			Where("user_id = ?", post.UserID).
			FirstOrCreate(&post).Error
		if err != nil {
			return err
		}
	}

	return nil
}
