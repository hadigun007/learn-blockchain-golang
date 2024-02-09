package main

import (
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    uint64
	transactions []string
}

func (b *Block) Print() {
	fmt.Printf("timestamp:		%d\n", b.timestamp)
	fmt.Printf("nonce:			%d\n", b.nonce)
	fmt.Printf("previousHash:		%s\n", b.previousHash)
	fmt.Printf("transactions:		%s\n", b.transactions)
}

func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		timestamp:    uint64(time.Now().UnixNano()),
	}
}

func main() {
	b := NewBlock(0, "init hash")
	b.Print()
}
