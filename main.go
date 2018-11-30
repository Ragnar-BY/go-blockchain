package main

import (
	"log"

	"github.com/Ragnar-BY/go-blockchain/types"
)

func main() {

	bc := types.NewBlockChain()

	cli := CLI{bc}
	log.Println(cli.Run())

	err := bc.CloseDB()
	if err != nil {
		log.Fatal(err)
	}

}
