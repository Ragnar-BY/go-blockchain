package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Ragnar-BY/go-blockchain/types"
)

type CLI struct {
	bc *types.Blockchain
}

func (cli *CLI) Run() error {
	blockData, printChain := cli.ParseFlags()
	if blockData != "" {
		err := cli.AddBlock([]byte(blockData))
		if err != nil {
			return err
		}
	}
	if printChain {
		cli.PrintBlockchain()
	}
	return nil
}

func (cli *CLI) ParseFlags() (string, bool) {
	var blockData string
	flag.StringVar(&blockData, "addblock", "", "Add new block with data")

	var printChain bool
	flag.BoolVar(&printChain, "printchain", false, " Print chain")
	flag.Parse()

	return blockData, printChain
}

func (cli *CLI) AddBlock(data []byte) error {
	return cli.bc.AddBlock(data)
}

func (cli *CLI) PrintBlockchain() {

	tip := cli.bc.Tip()

	block, err := cli.bc.GetBlockByHash(tip)
	if err != nil {
		log.Println(err)
	} else {
		for block != nil {
			fmt.Println(block.ToString())
			block, err = cli.bc.GetParentBlock(block)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
