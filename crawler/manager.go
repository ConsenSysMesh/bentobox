package crawler

import "sync"

type Manager struct {
	blocksInProcess BlocksInProcess
}

type BlocksInProcess struct {
	mutex sync.RWMutex
	Map   map[int]struct{}
	Cnt   int
}

func newManager() *Manager {
	response := &Manager{}
	response.blocksInProcess = BlocksInProcess{}
	response.blocksInProcess.Map = make(map[int]struct{})
	response.blocksInProcess.Cnt = 0

	return response
}

func (m *Manager) getBlocksInProcessCnt() int {
	m.blocksInProcess.mutex.Lock()
	defer m.blocksInProcess.mutex.Unlock()

	return m.blocksInProcess.Cnt
}

func (m *Manager) addBlockToProcess(id int) {
	m.blocksInProcess.mutex.Lock()
	defer m.blocksInProcess.mutex.Unlock()

	m.blocksInProcess.Map[id] = struct{}{}
	m.blocksInProcess.Cnt += 1
}

func (m *Manager) removeBlockToProcess(id int) {
	m.blocksInProcess.mutex.Lock()
	defer m.blocksInProcess.mutex.Unlock()

	delete(m.blocksInProcess.Map, id)
	m.blocksInProcess.Cnt -= 1
}
