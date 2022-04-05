package main

import (
	"fmt"
	"github.com/m0rphed/blockchain-go-brrr/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		// tips: if the data has changed
		//	-- we can determine the "data corruption" by comparing the hashes
		fmt.Println()
	}
}
