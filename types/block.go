package types

import (
	"bytes"
	"fmt"
	"time"

	"encoding/gob"
	"go-blockchain/utils"
)

type BlockHeader struct {
	PrevBlockHash [32]byte
	DataHash      [32]byte

	Time  int64
	Nonce [8]byte

	Hash [32]byte
}

func NewBlockHeader(prevBlockHash [32]byte, dataHash [32]byte) *BlockHeader {

	return &BlockHeader{PrevBlockHash: prevBlockHash, DataHash: dataHash}
}

func (bh *BlockHeader) FindNonce() {

	bh.Time = time.Now().UnixNano()
	nonce, hash := pow.Run(bh.HeaderNoNonce())

	bh.Hash = hash
	bh.Nonce = nonce

}

//check if blockHeader hash is under PoW target
func (bh *BlockHeader) Validate() bool {
	return pow.IsValid(bh.HeaderNoNonce(), bh.Nonce)
}

//prevBlockHash+dataHash
func (bh *BlockHeader) HeaderNoNonce() [32]byte {

	data := utils.Hash([]interface{}{
		bh.PrevBlockHash[:],
		bh.DataHash[:],
		bh.Time,
	})
	return data
}

type Block struct {
	Header *BlockHeader
	Data   []byte
	Number int64
}

func NewBlock(h *BlockHeader, data []byte, number int64) *Block {
	block := &Block{Header: h, Data: data, Number: number}
	return block
}

func (b *Block) Hash() [32]byte {

	return b.Header.Hash
}

//gob serialize and deserialize
func (b *Block) Serialize() ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	return result.Bytes(), err
}
func DeserializeBlock(b []byte) (*Block, error) {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&block)

	return &block, err
}

func (b *Block) ToString() string {
	t := time.Unix(0, b.Header.Time)
	str := fmt.Sprintf("Block %v:[PrevHash: %x, Data: [%s] , Hash %x, CreatedAt %v]",
		b.Number, b.Header.PrevBlockHash, b.Data, b.Hash(), t.Format("2006-01-02 15:04:05.99"))
	return str
}
