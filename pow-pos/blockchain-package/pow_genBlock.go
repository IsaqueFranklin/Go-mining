package blockchain

import (
	"strconv"
	"time"
)

//createGenesisBlock criar o primeiro bloco da blockchain.
func createGenesisBlock(difficulty int) Block {
	timestamp := time.Now().Unix()
	genesisBlock := Block{
		Index: 0,
		Timestamp: timestamp,
		PrevHash: "0",
		Data: "Genesis Block",
		Nonce: 0,
		Difficulty: difficulty,
	}

	genesisBlock.Hahs = calculateHash(genesisBlock)
	return genesisBlock
}