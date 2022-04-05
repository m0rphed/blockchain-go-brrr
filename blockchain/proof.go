package blockchain

/*
Consensus algorithm = proof of X algorithms

idea: secure the blockchain
		by forcing network to do work
		(work = computational power).

example: miners mining bitcoin = doing proof of work,
		so they can sign the block on the blockchain
	And the reason the get fees - it's because
	they are essentially powering the blockchain,
	and by doing it - they make the blocks and the data inside them more secure.

validation: So... when a user does computational work, it's necessary to provide a proof of work

idea:
	- the work must be hard to do
	- providing a proof must be relatively easy to do
*/

// steps:
//		take a data from the block
//		create a counter (nonce) which starts at 0
//		create a hash of the data plus the counter
//		check the hash to see if it meets a set of requirements
//		<repeat the last step if requirements was not met>

// requirements (=difficulty, which gets (could be) adjusted over the time):
//		-> First few bytes must contain 0s

import (
	"bytes"
	"math/big"
)

// Difficulty in this implementation would stay static, if we do a 'REAL' blockchain
//	we must somehow increase the difficulty over large period of time
//	-> main reason because:
//		- numbers of miners - increases,
//		- computational power - increases (also in general)
//	...and we need to make the time to mine a block stay the same, and block rates stay the same.
const Difficulty = 12

type ProofOfWork struct {
	Block *Block
	// Target is a pointer to a number
	// that represents the requirement,
	// which derive from difficulty
	Target *big.Int
}

// NewProof allows to take a pointer to a block
//	and produce a pointer to a proof of work
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// 256 is a number of bits in hash ('cause we use SHA256)
	// and then use the `target` to shift a number of bytes by N (=256-Difficulty)
	target.Lsh(target, uint(256-Difficulty)) // Lsh - is a left shift operation

	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte { // stopped 👉 7:16
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
		},
		[]byte{},
	)
	return data
}
