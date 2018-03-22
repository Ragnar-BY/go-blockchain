package types

import (
	//	"encoding/hex"
	"fmt"

	pow_ "github.com/Ragnar-BY/go-blockchain/pow"
	"github.com/Ragnar-BY/go-blockchain/utils"
)

var pow *pow_.ProofOfWork

type Blockchain3 struct {
	blocks []*Block
}

func NewBlockChain3() *Blockchain3 {
	pow = pow_.NewProofOfWork()
	return &Blockchain3{blocks: []*Block{GenesisBlock()}}
}
func (bc *Blockchain3) Blocks() []*Block {
	return bc.blocks
}

func (bc *Blockchain3) AddBlock(data []byte) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := prevBlock.Hash()

	dataHash := utils.Hash(data)
	header := NewBlockHeader(prevHash, dataHash)

	header.FindNonce()

	block := NewBlock(header, data)
	bc.blocks = append(bc.blocks, block)

	fmt.Println(block.ToString())
}

/*
const genPrevHash = "35353535353535353535353535353535353535353535353535353535353535353535353535353535353535353535353535353535353535353535353535353535"

func GenesisBlock() *Block {
	var prevHash [32]byte
	prevHashSlice, _ := hex.DecodeString(genPrevHash)
	copy(prevHash[:], prevHashSlice[:32])
	dataHash := utils.Hash([]byte{})

	header := NewBlockHeader(prevHash, dataHash)
	header.FindNonce()

	block := NewBlock(header, []byte{})
	return block
}
*/
