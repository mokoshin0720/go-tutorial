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

func gormConnect() *gorm.DB {
	dsn := "host=localhost user=postgres password=shinya dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", dsn)
	log.Println("postgresqlと接続")

	if err != nil {
		log.Fatal(err)
	}
	
	defer db.Close()
	return db
}