package crawler

import (
	"log"
	"time"

	gorp "gopkg.in/gorp.v1"
)

type Options struct {
	Host                 string
	MaxProcessingQueries int
	LoopTimeMs           int
}

func GetData(options Options, dbmap *gorp.DbMap) {
	var err error

	// Let's get the max height of this network
	maxBlockHeight, err := getNetworkHeight(options.Host)
	if err != nil {
		log.Fatalf("Error getting network height: %v", err)
	}

	// Start the manager
	m := newManager()

	// Inserting Loop
	// TODO

	// Feeding Loop
	for {
		// We get the ids of all the blocks we already have (ex: 1, 2, 5, 9)
		var obtainedIds []int
		_, err = dbmap.Select(&obtainedIds, "SELECT block_number_id FROM blocks ORDER BY block_number_id ASC")
		if err != nil {
			log.Printf("Error on query: %v", err)
			continue
		}
		obtainedIdsMap := make(map[int]struct{})
		for _, id := range obtainedIds {
			obtainedIdsMap[id] = struct{}{}
		}

		// How many blocks do we want to send to the processing queue?
		blocksToProcessCnt := options.MaxProcessingQueries - m.getBlocksInProcessCnt()
		blockIdsToQuery := make([]int, 0)
		for i := 0; i <= maxBlockHeight; i++ {
			if len(blockIdsToQuery) == blocksToProcessCnt {
				break
			}

			if _, ok := obtainedIdsMap[i]; ok {
				continue
			}

			blockIdsToQuery = append(blockIdsToQuery, i)
		}

		// We add the blockIds to the manager
		for _, id := range blockIdsToQuery {
			m.addBlockToProcess(id)
		}

		time.Sleep(time.Duration(options.LoopTimeMs) * time.Millisecond)
	}
}
