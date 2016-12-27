package main

import "github.com/consensys/event-crawler/database"

func main() {
	// Load SQL Database
	dbmap := database.InitDb()
	defer dbmap.Db.Close()
}
