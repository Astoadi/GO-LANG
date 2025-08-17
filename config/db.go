package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// var DB *sql.DB
// var DB *gorm.DB
var DB *sqlx.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}
	dsn := os.Getenv("DB_DSN")
	// db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, error := sql.Open("postgres", dsn)
	db, error := sqlx.Connect("postgres", dsn)
	if error != nil {
		fmt.Println("Error opening database: ", error)
		panic(error)
	}

	// sqlDb, err := db.DB()
	// if err != nil {
	// 	fmt.Println("Error opening database: ", error)
	// 	panic(err)
	// }
	// if err := sqlDb.Ping(); err != nil {
	// 	fmt.Println("Error connecting database: ", err)
	// 	panic(err)
	// }
	fmt.Println("Successfully connected to the database")
	DB = db
}
