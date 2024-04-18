package main

import (
	"fmt"
	"github.com/isaquefranklin/blockchain-consensus/blockchain"
)

func main() {
	difficulty := 3 //Ajuste conforme for necessário.

	//Testando proof of work.
	powBlockchain := []blockchain.Block{blockchain.CreateGenesisBlock(difficulty)}

	//Gere um novo bloco com dados usando proof of work.
	newBlockData := "Block Data."
	newPowBlock := blockchain.GenerateNewBlockWithPoW(powBlockchain[len(powBlockchain)-1], newBlockData, difficulty)
	powBlockchain = append(powBlockchain, newPowBlock)

	//Exibind a chain de proof of work.
	fmt.Println("Proof of Work Blockchain:")
	for _, block := range powBlockchain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Difficulty: %d\n", block.Difficulty)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}

	//Testando Proof of Stake
	posBlockchain := []blockchain.Block{blockchain.CreateGenesisBlockForPoS(difficulty)}

	// Gere um novo bloco com dados na Proof of Stake
	newPosBlock := blockchain.GenerateNewBlockWithPoS(posBlockchain[len(posBlockchain)-1], newBlockData, difficulty, []string{"Validator 1", "Validator 2", "Validator 3"})
	posBlockchain = append(posBlockchain, newPosBlock)

	//Exibindo a blockchain do proof of stake.
	fmt.Println("Proof of Stake Blockchain:")
	for _, block := range posBlockchain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Difficulty: %d\n", block.Difficulty)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}
