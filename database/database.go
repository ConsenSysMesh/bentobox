package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	gorp "gopkg.in/gorp.v1"
)

// TODO
// Feed these from a config file
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "mysecretpassword"
	DB_NAME     = "bentobox"
)

type Block struct {
	BlockNumber string `db:"block_number"`
	BlockHash   string `db:"block_hash"`
}

type Transaction struct {
	TransactionHash  string `db:"transaction_hash"`
	BlockNumber      string `db:"tx_block_number"`
	TransactionIndex string `db:"transaction_index"`
	From             string `db:"tx_from"`
	To               string `db:"tx_to"`
}

type Log struct {
	Id              int64  `db:"id"`
	TransactionHash string `db:"log_transaction_hash"`
	Data            string `db:"data"`
	LogIndex        string `db:"log_index"`
	Type            string `db:"mined"`
}

type Topic struct {
	LogId   int64  `db:"log_id"`
	Content string `db:"content"`
}

func InitDb() *gorp.DbMap {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalf("sql.Open failed %v", err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(Block{}, "blocks").SetKeys(true, "BlockNumber")
	dbmap.AddTableWithName(Transaction{}, "transactions").SetKeys(true, "TransactionHash")
	dbmap.AddTableWithName(Log{}, "logs").SetKeys(true, "Id")
	dbmap.AddTableWithName(Topic{}, "topics")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatalf("Create tables failed %v", err)
	}

	return dbmap
}
