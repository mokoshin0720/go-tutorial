// gormに依存するstruct
package model

import (
	"gorm.io/driver/postgres"
	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model // ID(プライマリキー), CreatedAt, UpdatedAt, DeletedAtをもつ
	Name string `gorm:"column:name"`
}

type Sample struct {
	Message string `json: "message"`
}

func GormConnect() *gorm.DB {
	dsn := "host=localhost user=shinya password=shinya dbname=user_signup port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	log.Println("postgresqlと接続")

	if err != nil {
		log.Fatal(err)
	}
	
	// defer db.Close()
	return db
}