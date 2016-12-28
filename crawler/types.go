package crawler

type ethBlockNumber struct {
	Result string `json:"result"`
}

type ethGetBlockByNumber struct {
	Result ethGetBlockByNumberData `json:"result"`
}

type ethGetBlockByNumberData struct {
	Hash   string                           `json:"hash"`
	Number string                           `json:"number"`
	Txs    []ethGetBlockByNumberTransaction `json:"transactions"`
}

type ethGetBlockByNumberTransaction struct {
	Hash string `json:"hash"`
}
