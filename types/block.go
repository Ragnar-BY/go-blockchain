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
	Nonce int64

	Hash [32]byte
}

func NewBlockHeader(prevBlockHash [32]byte, dataHash [32]byte) *BlockHeader {
	return &BlockHeader{PrevBlockHash: prevBlockHash, DataHash: dataHash}
}

func (bh *BlockHeader) FindNonce() {

	nonce, hash, t := pow.Run(bh.Header())
	bh.Hash = hash
	bh.Nonce = nonce
	bh.Time = t
}

//check if blockHeader hash is under PoW target
func (bh *BlockHeader) Validate() bool {
	return pow.IsValid(bh.FullHeader())
}

//prevBlockHash+dataHash
func (bh *BlockHeader) Header() []byte {
	data := bytes.Join(
		[][]byte{
			bh.PrevBlockHash[:], bh.DataHash[:],
		},
		[]byte{},
	)
	return data
}

//prevBlockHash+DataHash+time+nonce
func (bh *BlockHeader) FullHeader() []byte {

	timeByte := utils.IntToHex(bh.Time)

	data := bytes.Join(
		[][]byte{
			bh.Header(), timeByte,
		},
		[]byte{},
	)

	nonceByte := utils.IntToHex(bh.Nonce)
	data = bytes.Join(
		[][]byte{
			data, nonceByte,
		},
		[]byte{},
	)

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
