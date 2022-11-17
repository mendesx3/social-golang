package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitMySql() {
	var er error
	if er = godotenv.Load(); er != nil {
		log.Fatal(er)
	}

	url := os.Getenv("mysql.url")

	_, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("MySQL Connected!")
}
