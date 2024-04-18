package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

//O struct Block representa um bloco na blockchain (dãhh).
type Block struct {
	Index int
	Timestamp int64
	PrevHash string
	Data string
	Nonce int
	Difficulty int
	Hash string
}
