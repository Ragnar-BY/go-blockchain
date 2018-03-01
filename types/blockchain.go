package types

type Blockchain struct {
	blocks []*Block
}

func NewBlockChain() *Blockchain {
	return &Blockchain{blocks: []*Block{GenesisBlock()}}
}
func (bc *Blockchain) Blocks() []*Block {
	return bc.blocks
}

func (bc *Blockchain) AddBlock(data []byte) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := prevBlock.Hash()
	header := NewBlockHeader(prevHash, prevBlock.Number())
	block := NewBlock(header, data)
	bc.blocks = append(bc.blocks, block)
}
