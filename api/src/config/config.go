package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitMySql() (*sql.DB, error) {
	var er error
	if er = godotenv.Load(); er != nil {
		log.Fatal(er)
	}

	url := os.Getenv("mysql.url")

	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("MySQL Connected!")
	return db, nil
}
