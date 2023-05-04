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

func Init() {
	_, err := InitMySql()
	if err != nil {
		log.Fatal(err)
	}
}

func InitMySql() (*sql.DB, error) {
	var er error
	if er = godotenv.Load(); er != nil {
		log.Fatal(er)
	}
	url := os.Getenv("mysql.url")
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Println("Error on connecting to MySQL")
	}

	if err := db.Ping(); err != nil {
		db.Close()
		log.Println(`Error ping to MySQL`)
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("Connected to MySQL")
	return db, nil
}
