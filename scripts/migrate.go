package main

import (
	"database/sql"
	"github.com/deslee/cms/data"
)

func main() {
	db, err := sql.Open("sqlite3", "database.sqlite?_loc=auto")
	if err != nil {
		panic(err)
	}

	data.CreateTablesAndIndicesIfNotExist(db)
}
