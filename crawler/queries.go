package crawler

import (
	"log"
	"strconv"
)

func getNetworkHeight(url string) (response int, err error) {
	body := `{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":42}`
	target := ethBlockNumber{}
	if err = requestAndParseJSON(url, body, &target); err != nil {
		log.Printf("Query Error: %v", err)
		return -1, err
	}

	var r int64
	r, err = strconv.ParseInt(target.Result[2:], 16, 0)
	if err != nil {
		log.Printf("Parse Error: %v", err)
		return -1, err
	}

	response = int(r)
	return response, nil
}
