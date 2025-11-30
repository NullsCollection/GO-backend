package connection

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB
func ConnectDB() {
	godotenv.Load()

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        host, user, pass, dbname, port,
    )

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	
	DB = db
	fmt.Println("Connected to database")


}