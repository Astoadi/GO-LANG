package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// var DB *sql.DB
// var DB *gorm.DB
var DB *sqlx.DB
var MongoDB *mongo.Client

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
	fmt.Println("Successfully connected to the psql database")
	DB = db
}

func ConnectMongo() {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the mongodb database")
	MongoDB = client
}
