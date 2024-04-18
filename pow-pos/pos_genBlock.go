package blockchain

import (
	"strconv"
	"time"
)

func createGenesisBlockForPoS(difficulty int) Block {
	timestamp := time.Now().Unix()
	genesisBlock := Block{
	Index:        0,
	Timestamp:    timestamp,
	PrevHash:     "0",
	Data:         "Genesis Block",
	Nonce:        0,
	Difficulty:   difficulty,
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	return genesisBlock
}