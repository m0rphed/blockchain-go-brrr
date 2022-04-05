package blockchain

type BlockChain struct {
	Blocks []*Block
}

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
