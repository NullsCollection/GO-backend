package connection

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDB() {
	godotenv.Load()

	// Check if DATABASE_URL is provided 
	  databaseURL := os.Getenv("DATABASE_URL")

	  var dsn string
	  if databaseURL != "" {
		dsn = databaseURL
	  } else {
			host := os.Getenv("DB_HOST")
			user := os.Getenv("DB_USER")
			pass := os.Getenv("DB_PASS")
			dbname := os.Getenv("DB_NAME")
			port := os.Getenv("DB_PORT")
			sslmode := os.Getenv("DB_SSLMODE")

			 if sslmode == "" {
            sslmode = "disable"
        }

			dsn = fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        host, user, pass, dbname, port, sslmode,
    )

	  }


	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	
	DB = db
	fmt.Println("Connected to database")


}