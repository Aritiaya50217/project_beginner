package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Index     int       `json:"index"`
	Timestamp time.Time `json:"timestamp"`
	Symbol    string    `json:"symbol"`
	PrevHash  string    `json:"prev_hash"`
	Hash      string    `json:"hash"`
}

func CalculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp.String() + block.Symbol + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func NewBlock(index int, symbol, prevHash string) Block {
	block := Block{
		Index:     index,
		Timestamp: time.Now(),
		Symbol:    symbol,
		PrevHash:  prevHash,
	}
	block.Hash = CalculateHash(block)
	return block
}
