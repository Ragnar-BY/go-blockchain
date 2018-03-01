package types

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/crypto/sha3"
	"strconv"
	"time"
)

type BlockHeader struct {
	PrevBlockHash []byte
	Number        uint64
	Time          int64
}

func NewBlockHeader(prevBlockHash []byte, prevNumber uint64) *BlockHeader {
	return &BlockHeader{Time: time.Now().UnixNano(), PrevBlockHash: prevBlockHash, Number: prevNumber + 1}
}
func (bh *BlockHeader) Hash() []byte {

	t := []byte(strconv.FormatInt(bh.Time, 10))

	bNum := make([]byte, 8)
	binary.BigEndian.PutUint64(bNum, bh.Number)

	//prevHash+time+number
	b := append(t, bNum...)
	b = append(bh.PrevBlockHash[:], b...)

	hf := sha3.New256()
	hf.Write(b)

	return hf.Sum(nil)

}

type Block struct {
	header *BlockHeader
	data   []byte

	//cached
	hash []byte
}

func NewBlock(h *BlockHeader, data []byte) *Block {
	block := &Block{header: h, data: data}
	return block
}

func (b *Block) Hash() []byte {
	if len(b.hash) == 0 {
		b.hash = b.header.Hash()
	}
	return b.hash
}
func (b *Block) Data() []byte {
	return b.data
}
func (b *Block) Number() uint64 {
	return b.header.Number
}
func (b *Block) ToString() string {
	str := fmt.Sprintf("Block %v:[PrevHash: %x, Data: [%s] , Hash %x]",
		b.Number(), b.header.PrevBlockHash, b.data, b.Hash())
	return str
}

func GenesisBlock() *Block {
	header := NewBlockHeader([]byte{}, 0)
	return NewBlock(header, []byte{})
}
