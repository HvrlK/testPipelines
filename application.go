package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testPipelines/app"
)

func main() {
	db := app.TestDB()
	defer db.Close()

	server := app.NewServer(db)
	server.Start()
}