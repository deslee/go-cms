package main

import (
	"github.com/deslee/cms/data"
	"github.com/jinzhu/gorm"
	"log"
)
import _ "github.com/jinzhu/gorm/dialects/sqlite"

func main() {
	db, err := gorm.Open("sqlite3", "database.sqlite")
	die(err)
	db.LogMode(true)

	var user = data.User{}

	db.Debug().Where("Id = ?", "d8af4dd4-68e6-4707-a03e-b4e02af7675c").First(&user)

	log.Print("Hey!")

	defer db.Close()
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
