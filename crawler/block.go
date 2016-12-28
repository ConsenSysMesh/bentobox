package crawler

import (
	"log"
	"time"
)

func blockQuerying(m *Manager, options Options) {
	log.Printf("Starting Block Querying Loop")

	for {
		blockId := <-m.queueBlockChan

		go func() {
			block, blockTxHashes, err := getBlockData(options.Host, blockId)
			if err != nil {
				log.Printf("Error on query: %v", err)
			}

			// TODO
			// Use a library with levels of logs
			// Maybe the very ethereum's glog
			/*
				log.Printf("Got info for block %v (%v): hash: %v no of txs: %v",
					block.BlockNumber.String,
					block.BlockNumberId,
					block.BlockHash.String,
					len(blockTxHashes),
				)
			*/

			m.insertBlock(block)

			// TODO
			// Insert tx hashes!
			_ = blockTxHashes
		}()
	}
}

func blockInsertingToDB(m *Manager) {
	log.Println("Starting Block Inserting to DB Loop")

	for {
		block := <-m.insertBlockToDBChan
		if err := m.dbmap.Insert(block); err != nil {
			log.Printf("Error inserting block %v: %v", block.BlockNumberId, err)
		}
	}
}

func blockQueueFeeding(m *Manager, options Options) {
	log.Println("Starting Block Queue Feeding Loop")

	for {
		// Let's get the current max height of this network
		maxBlockHeight, err := getNetworkHeight(options.Host)
		if err != nil {
			log.Fatalf("Error getting network height: %v", err)
		}
		log.Printf("Block Height %v", maxBlockHeight)

		// We get the ids of all the blocks we already have (ex: 1, 2, 5, 9)
		var obtainedIds []int64
		_, err = m.dbmap.Select(&obtainedIds, "SELECT block_number_id FROM blocks ORDER BY block_number_id ASC")
		if err != nil {
			log.Printf("Error on query: %v", err)
			continue
		}
		obtainedIdsMap := make(map[int64]struct{})
		for _, id := range obtainedIds {
			obtainedIdsMap[id] = struct{}{}
		}
		log.Printf("Got %v block ids from the database", len(obtainedIds))

		// How many blocks do we want to send to the processing queue?
		blocksToQueueCount := options.MaxProcessingQueries - m.getBlocksInProcessCnt()
		blockIdsToQuery := make([]int64, 0)
		var i int64
		for i = 0; i <= maxBlockHeight; i++ {
			if len(blockIdsToQuery) == blocksToQueueCount {
				break
			}

			if _, ok := obtainedIdsMap[i]; ok {
				continue
			}

			blockIdsToQuery = append(blockIdsToQuery, i)
		}

		// We add the blockIds to the manager
		log.Printf("Will add %v block Ids to the manager to query", len(blockIdsToQuery))
		for _, id := range blockIdsToQuery {
			m.addBlockToProcess(id)
		}

		time.Sleep(time.Duration(options.LoopTimeMs) * time.Millisecond)
	}
}
