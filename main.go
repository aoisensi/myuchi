package main

import (
	"flag"
	"log"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/go-martini/martini"
)

var (
	optDBPath              string
	optCFGName             string
	dbUser, dbRoom, dbPost *db.Col
)

func init() {
	flag.StringVar(&optDBPath, "db", "./db", "database path")
	flag.StringVar(&optCFGName, "cfg", "./config.json", "config json file name")
	initDB()
	initConfig()
	initVerify()
}

func main() {
	log.Println("Start Server")
	flag.Parse()
	m := martini.Classic()
	m.Get("/verify/:type", verify)
	m.Get("/verify/:type/redirect")
	log.Println("Start Server")
	m.Run()
}
