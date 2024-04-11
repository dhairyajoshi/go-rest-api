package database

import (
	"database/sql"
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var once sync.Once

var connection *sql.DB

func GetPostgresConnection() *sql.DB {

	if connection != nil {
		return connection
	}

	once.Do(func() {
		connStr := "postgres://postgres:pass@localhost/test?sslmode=disable"
		db, _ := sql.Open("postgres", connStr)
		connection = db
	})

	err := connection.Ping()

	if err != nil {
		var msg string = fmt.Sprintf("couldn't connect to database: %v", err)
		panic(msg)
	}

	println("database connection successful!")

	return connection
}

func GetDbConnection() *gorm.DB {
	conn := GetPostgresConnection()

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{})

	if err != nil {
		errstr := fmt.Sprintf("error connecting to db: %v", err)
		panic(errstr)
	}

	return db
}
