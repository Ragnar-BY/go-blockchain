package types

import (
	"encoding/hex"
	"fmt"

	pow_ "go-blockchain/pow"
	"go-blockchain/utils"
)

var pow *pow_.ProofOfWork

type Blockchain struct {
	blocks []*Block
}

func NewBlockChain() *Blockchain {
	pow = pow_.NewProofOfWork()
	return &Blockchain{blocks: []*Block{GenesisBlock()}}
}
func (bc *Blockchain) Blocks() []*Block {
	return bc.blocks
}

func (bc *Blockchain) AddBlock(data []byte) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := prevBlock.Hash()

	dataHash := utils.Hash(data)
	header := NewBlockHeader(prevHash, dataHash)

	header.FindNonce()

	block := NewBlock(header, data, prevBlock.Number()+1)
	bc.blocks = append(bc.blocks, block)

	fmt.Println(block.ToString())
}

const genPrevHash = "35353535353535353535353535353535353535353535353535353535353535353535353535353535353535353535353535353535353535353535353535353535"

func GenesisBlock() *Block {
	var prevHash [32]byte
	prevHashSlice, _ := hex.DecodeString(genPrevHash)
	copy(prevHash[:], prevHashSlice[:32])
	dataHash := utils.Hash([]byte{})

	header := NewBlockHeader(prevHash, dataHash)
	header.FindNonce()

	block := NewBlock(header, []byte{}, 0)
	return block
}
