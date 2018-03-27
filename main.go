package main

import (
	"github.com/Ragnar-BY/go-blockchain/types"
)

func main() {

	bc := types.NewBlockChain()
	defer bc.CloseDB()

	cli := CLI{bc}
	cli.Run()

}
