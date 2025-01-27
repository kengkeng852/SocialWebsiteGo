package main

import (
	"github.com/kengkeng852/socialwebsitego/config"
	"github.com/kengkeng852/socialwebsitego/databases"
	"github.com/kengkeng852/socialwebsitego/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	tx := db.Connect().Begin()

	commentMigration(tx)
	followMigration(tx)
	postMigration(tx)
	userMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}

}

func commentMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Comment{})
}

func followMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Follow{})
}

func postMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Post{})
}

func userMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.User{})
}
 