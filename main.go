package main

import (
	"fmt"
	"strings"
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
	fmt.Printf("transactions:		%s\n\n", b.transactions)
}

func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		timestamp:    uint64(time.Now().UnixNano()),
	}
}

type Blockchain struct {
	transacrionPool []string
	chain           []*Block
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf(strings.Repeat("*", 25))
}

func main() {
	blockchain := NewBlockchain()
	blockchain.CreateBlock(1, "Hash 1")
	blockchain.CreateBlock(2, "Hash 2")
	blockchain.Print()
}
