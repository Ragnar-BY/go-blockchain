package main

import (
	"log"

	"github.com/Ragnar-BY/go-blockchain/types"
)

func main() {

	bc := types.NewBlockChain()
	defer bc.CloseDB()

	cli := CLI{bc}
	log.Println(cli.Run())

}
