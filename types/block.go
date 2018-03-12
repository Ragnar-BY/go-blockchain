package types

import (
	"bytes"
	"fmt"
	"time"

	"go-blockchain/utils"
)

type BlockHeader struct {
	prevBlockHash [32]byte
	dataHash      [32]byte

	time  int64
	nonce int64

	hash [32]byte
}

func NewBlockHeader(prevBlockHash [32]byte, dataHash [32]byte) *BlockHeader {
	return &BlockHeader{prevBlockHash: prevBlockHash, dataHash: dataHash}
}

func (bh *BlockHeader) FindNonce() {

	nonce, hash, t := pow.Run(bh.Header())
	bh.hash = hash
	bh.nonce = nonce
	bh.time = t
}

//check if blockHeader hash is under PoW target
func (bh *BlockHeader) Validate() bool {
	return pow.IsValid(bh.FullHeader())
}

func (bh *BlockHeader) Hash() [32]byte {
	return bh.hash
}

//prevBlockHash+dataHash
func (bh *BlockHeader) Header() []byte {
	data := bytes.Join(
		[][]byte{
			bh.prevBlockHash[:], bh.dataHash[:],
		},
		[]byte{},
	)
	return data
}

//prevBlockHash+DataHash+time+nonce
func (bh *BlockHeader) FullHeader() []byte {

	timeByte := utils.IntToHex(bh.time)

	data := bytes.Join(
		[][]byte{
			bh.Header(), timeByte,
		},
		[]byte{},
	)

	nonceByte := utils.IntToHex(bh.nonce)
	data = bytes.Join(
		[][]byte{
			data, nonceByte,
		},
		[]byte{},
	)

	return data
}

type Block struct {
	header *BlockHeader
	data   []byte
	number int64
}

func NewBlock(h *BlockHeader, data []byte, number int64) *Block {
	block := &Block{header: h, data: data, number: number}
	return block
}

func (b *Block) Hash() [32]byte {

	return b.header.hash
}
func (b *Block) Data() []byte {
	return b.data
}
func (b *Block) Number() int64 {
	return b.number
}
func (b *Block) ToString() string {
	t := time.Unix(0, b.header.time)
	str := fmt.Sprintf("Block %v:[PrevHash: %x, Data: [%s] , Hash %x, CreatedAt %v]",
		b.Number(), b.header.prevBlockHash, b.data, b.Hash(), t.Format("2006-01-02 15:04:05.99"))
	return str
}
