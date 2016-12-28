package crawler

import (
	"sync"

	"github.com/consensys/bentobox-crawler/database"

	gorp "gopkg.in/gorp.v1"
)

type Manager struct {
	dbmap               *gorp.DbMap
	blocksInQueue       BlocksInQueue
	queueBlockChan      chan int64
	insertBlockToDBChan chan *database.Block
}

type BlocksInQueue struct {
	mutex sync.RWMutex
	items map[int64]struct{}
	count int
}

func newManager() *Manager {
	return &Manager{
		blocksInQueue: BlocksInQueue{
			items: make(map[int64]struct{}),
			count: 0,
		},
		queueBlockChan:      make(chan int64),
		insertBlockToDBChan: make(chan *database.Block),
	}
}

func (m *Manager) getBlocksInProcessCnt() int {
	m.blocksInQueue.mutex.Lock()
	defer m.blocksInQueue.mutex.Unlock()
	return m.blocksInQueue.count
}

func (m *Manager) addBlockToProcess(id int64) {
	m.blocksInQueue.mutex.Lock()
	defer m.blocksInQueue.mutex.Unlock()
	m.blocksInQueue.items[id] = struct{}{}
	m.blocksInQueue.count += 1
	m.queueBlockChan <- id
}

func (m *Manager) removeBlockToProcess(id int64) {
	m.blocksInQueue.mutex.Lock()
	defer m.blocksInQueue.mutex.Unlock()
	delete(m.blocksInQueue.items, id)
	m.blocksInQueue.count -= 1
}

func (m *Manager) insertBlock(block *database.Block) {
	m.insertBlockToDBChan <- block
	m.removeBlockToProcess(block.BlockNumberId)
}
