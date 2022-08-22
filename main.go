package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// creates array of Block Addresses
type BlockChain struct {
	blocks []*Block
}

// timestamp, block height etc are added when system becomes complicated
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// method of Block struct
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	// used [:] to create slice from an array
	b.Hash = hash[:]
}

// Block is created on address
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// method of BlockChain struct
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// returns pointer to the Block created by CreateBlock
// using Genesis as the data and empty previous hash value
// Genesis means "the origin or mode of formation of something"
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// returns pointer addresses to the BlockChain struct
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
