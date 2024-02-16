package initializers

import (
	"log"
	"main/db"
	"os"
)

func InitializeDB() {
	dsn := os.Getenv("DSN")
	err := db.InitDB(dsn)

	if err != nil {
		log.Fatalln(err)
	}
}
