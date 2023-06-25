package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte // it's sort of back-linked list
	Nonce    int
}

// CreateBlock allows to create a block
// based on the hash of prev block and some data
func CreateBlock(data string, prevHash []byte) *Block {
	// initial blocks nonce is 0
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Genesis adds "genesis block"
// - because the impl. requires a first block to always exist in the chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Serialize block in the blockchain
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)

	if err != nil {
		log.Panic(err)
	}

	return res.Bytes() // byte representation of the block
}

// Deserialize data into block
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)

	if err != nil {
		log.Panic(err)
	}

	return &block // reference to the block which we created inside this function
}
