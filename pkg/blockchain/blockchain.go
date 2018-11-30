package blockchain

import (
	"encoding/hex"

	pow_ "github.com/Ragnar-BY/go-blockchain/pkg/blockchain/pow"
	"github.com/Ragnar-BY/go-blockchain/pkg/utils"
)

var pow *pow_.ProofOfWork

type blockDatabase interface {
	AddNewBlock(hash [32]byte, serial []byte) error
	GetLastHash() ([32]byte, error)
	GetBlockByHash(hash [32]byte) []byte
	CreateIfNotExist() (bool, error)
}

// Blockchain is struct for blockchain
type Blockchain struct {
	tip [32]byte
	db  blockDatabase
}

// Tip returns tip.
func (bc *Blockchain) Tip() [32]byte {
	return bc.tip
}

// NewBlockChain create blockchain.
func NewBlockChain(db blockDatabase) (*Blockchain, error) {

	pow = pow_.NewProofOfWork()

	bc := Blockchain{db: db}
	exist, err := bc.db.CreateIfNotExist()
	if err != nil {
		return nil, err
	}
	if !exist {
		genesis, err := GenesisBlock()
		if err != nil {
			return nil, err
		}
		serial, err := genesis.Serialize()
		if err != nil {
			return nil, err
		}
		hash := genesis.Hash()

		err = bc.db.AddNewBlock(hash, serial)
		if err != nil {
			return nil, err
		}
		bc.tip = hash
	} else {
		tip, err := bc.db.GetLastHash()
		if err != nil {
			return nil, err
		}
		bc.tip = tip
	}
	return &bc, nil

}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data []byte) error {
	var lastHash [32]byte

	lastHash, err := bc.db.GetLastHash()
	if err != nil {
		return err
	}

	dataHash, err := utils.Hash(data)
	if err != nil {
		return err
	}
	header := NewBlockHeader(lastHash, dataHash)

	err = header.FindNonce()
	if err != nil {
		return err
	}

	newBlock := NewBlock(header, data)

	serial, err := newBlock.Serialize()
	if err != nil {
		return err
	}

	err = bc.db.AddNewBlock(newBlock.Hash(), serial)
	if err != nil {
		return err
	}
	bc.tip = newBlock.Hash()
	return nil
}

// GetBlockByHash gets block by hash
func (bc *Blockchain) GetBlockByHash(hash [32]byte) (*Block, error) {
	serial := bc.db.GetBlockByHash(hash)
	if serial == nil {
		return nil, nil
	}
	b, err := DeserializeBlock(serial)
	return b, err
}

// GetParentBlock gets parent block for block.
func (bc *Blockchain) GetParentBlock(block *Block) (*Block, error) {
	prevHash := block.Header.PrevBlockHash
	b, err := bc.GetBlockByHash(prevHash)
	return b, err
}

const genPrevHash = "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001"

// GenesisBlock generates genesis block.
func GenesisBlock() (*Block, error) {
	var prevHash [32]byte
	prevHashSlice, _ := hex.DecodeString(genPrevHash)
	copy(prevHash[:], prevHashSlice[:32])
	dataHash, err := utils.Hash([]byte{})
	if err != nil {
		return nil, err
	}
	header := NewBlockHeader(prevHash, dataHash)
	err = header.FindNonce()
	if err != nil {
		return nil, err
	}

	block := NewBlock(header, []byte{})
	return block, nil
}
