package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Connect : Instantiate db connection
func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("ppi.db"), &gorm.Config{})
	handle(err)

	db.AutoMigrate(&Project{})

	return db

}
