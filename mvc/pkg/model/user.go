// gormに依存するstruct

package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"time"
)

type User struct {
	ID        int       `gorm:"primary_key" "column:id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:createdat"`
}

func main() {
	db, err := sql.Open("mysql",  "root@/sample?charset=utf8&parseTime=True&loc=Local")
	log.Println("mysqlと接続")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}