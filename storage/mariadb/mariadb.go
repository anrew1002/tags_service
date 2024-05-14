package mariadb

import (
	// "cms-golang-service/pkg"

	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"isustud.com/m/internal/config"
)

func New(cfg config.Config) *sqlx.DB {
	user := "tag_service"
	dbName := "tag_service"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci", user, cfg.DBPassword, cfg.DBHost, dbName)
	fmt.Println(connectionString)
	db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Failed set connection to DB, ", err)
	}
	db.SetConnMaxIdleTime(time.Minute * 15)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err = db.Ping(); err != nil {
		log.Fatal("Failed ping DB ", err)
	}

	return db
}
