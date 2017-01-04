package crawler

import (
	"sync"

	"github.com/consensys/bentobox-crawler/database"

	gorp "gopkg.in/gorp.v1"
)

type Manager struct {
	dbmap               *gorp.DbMap
	options             Options
	blocksQueue         BlocksQueue
	queryBlockToRPCChan chan int64
	insertBlockToDBChan chan *database.Block
	insertTxToDBChan    chan *database.Transaction
}

type BlocksQueue struct {
	mutex sync.RWMutex
	items map[int64]struct{}
	count int
}

type Options struct {
	Host                 string
	MaxProcessingQueries int
	LoopTimeMs           int
}

func newManager(options Options, dbmap *gorp.DbMap) *Manager {
	return &Manager{
		options: options,
		dbmap:   dbmap,
		blocksQueue: BlocksQueue{
			items: make(map[int64]struct{}),
			count: 0,
		},
		queryBlockToRPCChan: make(chan int64),
		insertBlockToDBChan: make(chan *database.Block),
		insertTxToDBChan:    make(chan *database.Transaction),
	}
}

func (m *Manager) addBlockToQueue(id int64) {
	m.blocksQueue.mutex.Lock()
	defer m.blocksQueue.mutex.Unlock()

	if _, ok := m.blocksQueue.items[id]; !ok {
		m.blocksQueue.items[id] = struct{}{}
		m.blocksQueue.count += 1
		m.queryBlockToRPCChan <- id
	}
}

func (m *Manager) removeBlockFromQueue(id int64) {
	m.blocksQueue.mutex.Lock()
	defer m.blocksQueue.mutex.Unlock()

	delete(m.blocksQueue.items, id)
	m.blocksQueue.count -= 1
}

func (m *Manager) getBlocksQueueCount() int {
	m.blocksQueue.mutex.Lock()
	defer m.blocksQueue.mutex.Unlock()

	return m.blocksQueue.count
}
