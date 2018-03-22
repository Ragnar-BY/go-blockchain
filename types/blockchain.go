package types

import (
	"encoding/hex"
	"log"

	pow_ "github.com/Ragnar-BY/go-blockchain/pow"
	"github.com/Ragnar-BY/go-blockchain/utils"
)

var pow *pow_.ProofOfWork

const dbFile = "blockchain.db"

type Blockchain struct {
	tip [32]byte
	db  *utils.Database
}

func NewBlockChain() *Blockchain {

	pow = pow_.NewProofOfWork()

	db, err := utils.OpenDB(dbFile)
	if err != nil {
		log.Panic(err)
	}
	bc := Blockchain{db: db}

	if bc.db.IsBucketExist() == false {

		bc.db.CreateNewBucket()

		genesis := GenesisBlock()
		serial, err := genesis.Serialize()
		hash := genesis.Hash()
		if err != nil {
			log.Panic(err)
		}

		err = bc.db.AddNewBlock(hash, serial)
		if err != nil {
			log.Panic(err)
		}
		bc.tip = hash
	} else {
		tip, err := bc.db.GetLastHash()
		if err != nil {
			log.Panic(err)
		}
		bc.tip = tip
	}
	return &bc

}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data []byte) {
	var lastHash [32]byte

	lastHash, err := bc.db.GetLastHash()
	if err != nil {
		log.Panic(err)
	}

	dataHash := utils.Hash(data)
	header := NewBlockHeader(lastHash, dataHash)

	header.FindNonce()

	newBlock := NewBlock(header, data)

	serial, err := newBlock.Serialize()
	if err != nil {
		log.Panic(err)
	}

	err = bc.db.AddNewBlock(newBlock.Hash(), serial)
	if err != nil {
		log.Panic(err)
	}

	log.Println(newBlock.ToString())
}

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
