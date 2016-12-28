package crawler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/consensys/bentobox-crawler/database"
)

func getNetworkHeight(url string) (response int64, err error) {
	body := `{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":42}`
	target := ethBlockNumber{}
	if err = requestAndParseJSON(url, body, &target); err != nil {
		log.Printf("Query Error: %v", err)
		return -1, err
	}

	response, err = strconv.ParseInt(target.Result[2:], 16, 0)
	if err != nil {
		log.Printf("Parse Error: %v", err)
		return -1, err
	}

	return
}

func getBlockData(url string, id int64) (block *database.Block, blockTxHashes []string, err error) {
	hexID := fmt.Sprintf("0x%x", id)
	body := `{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["` + hexID + `", true],"id":42}`
	target := ethGetBlockByNumber{}
	if err = requestAndParseJSON(url, body, &target); err != nil {
		log.Printf("Query Error: %v", err)
		return
	}

	block = &database.Block{}
	block.BlockNumberId = id
	block.BlockNumber.Scan(target.Result.Number)
	block.BlockHash.Scan(target.Result.Hash)
	for _, tx := range target.Result.Txs {
		blockTxHashes = append(blockTxHashes, tx.Hash)
	}

	return
}
