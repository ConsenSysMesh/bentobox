package crawler

import gorp "gopkg.in/gorp.v1"

type Options struct {
	Host                 string
	MaxProcessingQueries int
	LoopTimeMs           int
}

func GetData(options Options, dbmap *gorp.DbMap) {
	m := newManager()
	m.dbmap = dbmap

	go blockQuerying(m, options)
	go blockInsertingToDB(m)
	go blockQueueFeeding(m, options)
}
