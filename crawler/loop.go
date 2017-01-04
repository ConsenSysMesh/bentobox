package crawler

import gorp "gopkg.in/gorp.v1"

func GetData(options Options, dbmap *gorp.DbMap) {
	manager := newManager(options, dbmap)

	go manager.insertBlockPipe()
	go manager.queryBlockDispatcher()
	go manager.feedBlocksToQueue()
}
