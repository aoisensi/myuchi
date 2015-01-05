package main

import (
	"log"

	"github.com/HouzuoGuo/tiedot/db"
)

func initDB() {

	db, err := db.OpenDB(optDBPath)
	if err != nil {
		log.Panicln(err.Error())
	}
	dbUser = createOrLoad(db, "user")
	dbRoom = createOrLoad(db, "room")
	dbPost = createOrLoad(db, "post")
	return
}

func createOrLoad(db *db.DB, name string) *db.Col {
	col := db.Use(name)
	if col == nil {
		db.Create(name)
		col = db.Use(name)
	}
	return col
}
