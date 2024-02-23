package finalizers

import (
	"log"
	"main/db"
)

func ShutdownDB() {
	err := db.CloseDB()
	if err != nil {
		log.Fatalln("Error closing the database connection:", err)
	}
}
