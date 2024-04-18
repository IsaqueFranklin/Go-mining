package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

//Para calcular o Hash de um bloco primeiro é necessário concatenar todos seus campos e aplicar a função Hash.
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + strconv.FormatInt(block.Timestamp, 10) + block.PrevHash + block.Data + strconv.Itoa(block.Nonce)

	hash := sha256.Sum256([]byte(record))

	return hex.EncodeToString(hash[:])
}