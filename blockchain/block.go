package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte // it's sort of back-linked list
}

// DeriveHash allows
// to create the hash based on the previous hash and the data
func (b *Block) DeriveHash() {
	// join data of current hash and combine with prev bytes
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info) // actual hash
	b.Hash = hash[:]
}

// CreateBlock allows to create a block
// based on the hash of prev block and some data
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// AddBlock adds new block to blockchain with specified data
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis adds "genesis block"
// - because the impl. requires a first block to always exist in the chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
