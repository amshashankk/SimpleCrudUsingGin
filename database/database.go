package database

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var ConnectDB *gorm.DB
var err error

func Connect() {
	ConnectDB, err = gorm.Open("postgres", "user=postgres dbname=Gin password=mobcoder@123 sslmode=disable")
	if err != nil {
		panic(err)
	}
}