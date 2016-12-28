package main

import (
	"flag"

	"github.com/consensys/bentobox-crawler/crawler"
	"github.com/consensys/bentobox-crawler/database"
)

func main() {
	// Process startup parameters
	optionsDBUser := flag.String("dbuser", "postgres", "database username")
	optionsDBPassword := flag.String("dbpassword", "mysecretpassword", "database password")
	optionsDBName := flag.String("dbname", "bentobox", "database name")
	optionsHost := flag.String("host", "http://127.0.0.1:8545", "URL of the ethereum node RPC")
	optionsMaxProcessingQueries := flag.Int("maxprocessingqueries", 100,
		"Maximum of concurrent queries to the RPC Node")
	optionsLoopTimeMs := flag.Int("looptimems", 5000, "Iteration interval for block querying")
	flag.Parse()

	// Load SQL Database
	dbOpts := database.Options{
		User:     *optionsDBUser,
		Password: *optionsDBPassword,
		DBName:   *optionsDBName,
	}
	dbmap := database.InitDb(dbOpts)
	defer dbmap.Db.Close()

	// Start feeding loop
	crawlerOpts := crawler.Options{
		Host:                 *optionsHost,
		MaxProcessingQueries: *optionsMaxProcessingQueries,
		LoopTimeMs:           *optionsLoopTimeMs,
	}
	crawler.GetData(crawlerOpts, dbmap)

	// There you are!
	select {}
}
