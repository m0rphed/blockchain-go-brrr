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
//		- take a data from the block
//		- create a counter (nonce) which starts at 0
//		- create a hash of the data plus the counter
//		- check the hash to see if it meets a set of requirements
//		<repeat the last step if requirements was not met>

// requirements (=difficulty, which gets (could be) adjusted over the time):
//		-> First few bytes must contain 0s

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Difficulty in this implementation would stay static, if we do a 'REAL' blockchain
//	we must somehow increase the difficulty over large period of time
//	-> main reason because:
//		- numbers of miners - increases,
//		- computational power - increases (also in general)
//	...and we need to make the time to mine a block stay the same, and block rates stay the same.
const Difficulty = 18

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

// InitData helps combine data with the nonce
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			// cast nonce and difficulty
			// to int64, then convert to bytes
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0
	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce) // prepare data
		hash = sha256.Sum256(data)  // hash the data

		fmt.Printf("\r%x", hash) // printing: see the result
		// convert hash to big integer
		intHash.SetBytes(hash[:])

		// compare the new hash to the PoW target
		if intHash.Cmp(pow.Target) == -1 {
			// if hash is smaller than target
			// then we have a valid hash
			// (means: we can sign the block)
			// and we can return it
			return nonce, hash[:]
		} else {
			nonce++
		}
	}
	fmt.Println() // printing: separator
	return nonce, hash[:]
}

/*
 `Validate` function idea:
 	after we run the PoW `Run` func. - we have the nonce,
	which will allow us to derive
	the hash which met the `target` we wanted,
	and then we'll be able to run that cycle again
	to show that the hash is valid.

	=> security: based on idea that you want change a block inside the chain
		-- you'll have to recalculate the hash itself (large amount of time),
		and then recalculate each block's hash (after that block) as well to validate the data.

		So, the actual work to create/sign a block is very difficult/time-consuming,
		when validation is pretty easy.
*/

// Validate validates the proof of work
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)

	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target) == -1
}

// ToHex converts a number to a byte array
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	// organize the data in big endian notation
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
