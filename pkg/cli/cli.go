package cli

import (
	"flag"
	"fmt"

	"github.com/Ragnar-BY/go-blockchain/pkg/blockchain"
)

// CLI is command-line interface
type CLI struct {
	Blockchain *blockchain.Blockchain
}

// Run runs program
func (cli *CLI) Run() error {
	blockData, printChain := cli.parseFlags()
	if blockData != "" {
		err := cli.addBlock([]byte(blockData))
		if err != nil {
			return err
		}
	}
	if printChain {
		s, err := cli.printBlockchain()
		if err != nil {
			return err
		}
		fmt.Println(s)
	}
	return nil
}

func (cli *CLI) parseFlags() (string, bool) {
	var blockData string
	flag.StringVar(&blockData, "addblock", "", "Add new block with data")

	var printChain bool
	flag.BoolVar(&printChain, "printchain", false, " Print chain")
	flag.Parse()

	return blockData, printChain
}

func (cli *CLI) addBlock(data []byte) error {
	return cli.Blockchain.AddBlock(data)
}

func (cli *CLI) printBlockchain() (string, error) {

	tip := cli.Blockchain.Tip()
	blockStr := ""
	block, err := cli.Blockchain.GetBlockByHash(tip)
	if err != nil {
		return "", err
	}
	for block != nil {
		blockStr += block.ToString() + "\n"
		block, err = cli.Blockchain.GetParentBlock(block)
		if err != nil {
			return "", err
		}
	}

	return blockStr, err
}
